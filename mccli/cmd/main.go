// Licensed Materials - Property of IBM
// (C) Copyright IBM Corp. 2019. All Rights Reserved.
// US Government Users Restricted Rights - Use, duplication or
// disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
// Copyright 2019 IBM Corporation

package main

import (
	"flag"
	"log"
	"os"

	// This next line lets us use IBM Kubernetes Service
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"

	mcCliPkg "github.ibm.com/istio-research/mc2019/mccli/pkg"
)

func main() {
	var namespace string
	flag.StringVar(&namespace, "namespace", "", "Kubernetes namespace")
	var kcontext string
	flag.StringVar(&kcontext, "context", "", "Kubernetes configuration context")
	flag.Parse()

	var ns string
	if namespace != "" {
		ns = namespace
	} else {
		ns = "default"
	}

	cl, err := mcCliPkg.NewCliClient(namespace, kcontext)
	if err != nil {
		log.Fatalf("Failed to create cl: %s", err)
	}
	expositions, err := mcCliPkg.GetExposures(cl, ns)
	if err != nil {
		log.Fatalf("Failed to list exposures: %s", err)
	}
	// log.Printf("%d expositions in %q\n", len(*expositions), ns)
	// for _, exposition := range *expositions {
	// 	log.Printf("Exposition: %q\n", exposition.ObjectMeta.GetName())
	// }

	openAPI, err := mcCliPkg.Convert(cl, *expositions)
	if err != nil {
		log.Fatalf("Failed to convert: %s", err)
	}

	_ = mcCliPkg.ToYAML(openAPI, os.Stdout)
}