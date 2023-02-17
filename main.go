package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	values "helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

func main() {
	options := values.Options{
		ValueFiles: os.Args[1:],
	}
	merged, err := options.MergeValues([]getter.Provider{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "err %v\n", err)
		return
	}
	out, err := yaml.Marshal(merged)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err %v\n", err)
		return
	}
	fmt.Printf("%v", string(out))
}
