package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/Andriizachepilo/cicd-golang/cicd" 
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: cicd.Provider,
    })
}
