package main

import (
    "os/exec"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "fmt"
)

func resourceBuild() *schema.Resource {
    return &schema.Resource{
        Create: resourceBuildCreate,
        Read:   resourceBuildRead,
        Delete: resourceBuildDelete,

        Schema: map[string]*schema.Schema{
            "programming_language": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceBuildCreate(d *schema.ResourceData, m interface{}) error {
    // Execute Maven build command
    cmd := exec.Command("mvn", "clean", "package")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error executing Maven build command: %s, output: %s", err, string(output))
    }

    // Log success message
    fmt.Printf("Maven build successful: %s\n", string(output))

    // Set the resource ID
    d.SetId("build_step")
    return nil
}

func resourceBuildRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceBuildDelete(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}
