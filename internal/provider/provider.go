// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package virtualbox serves as an entrypoint, returning the list of available
// resources for the plugin.
package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terra-farm/go-virtualbox"
)

func init() {
	// Terraform is already adding the timestamp for us
	log.SetFlags(log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("pid-%d-", os.Getpid()))
}

// ProviderConfig holds the provider configuration
type ProviderConfig struct {
	Manager       *virtualbox.Manager
	MachineFolder string
	GoldFolder    string
}

// New returns a resource provider for virtualbox.
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"machine_folder": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the folder where VirtualBox machines will be stored. Defaults to ~/.terraform/virtualbox/machine",
			},
			"gold_folder": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the folder where gold images will be stored. Defaults to ~/.terraform/virtualbox/gold",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"virtualbox_vm": resourceVM(),
		},
		ConfigureContextFunc: configure,
	}
}

// configure creates a new instance of the new virtualbox manager which will be
// used for communication with virtualbox.
func configure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	manager := virtualbox.NewManager()
	
	// Get machine folder from config or use default
	machineFolder := d.Get("machine_folder").(string)
	if machineFolder == "" {
		usr, err := os.UserHomeDir()
		if err != nil {
			return nil, diag.Errorf("unable to get the current user home directory: %v", err)
		}
		machineFolder = filepath.Join(usr, ".terraform/virtualbox/machine")
	}
	
	// Get gold folder from config or use default
	goldFolder := d.Get("gold_folder").(string)
	if goldFolder == "" {
		usr, err := os.UserHomeDir()
		if err != nil {
			return nil, diag.Errorf("unable to get the current user home directory: %v", err)
		}
		goldFolder = filepath.Join(usr, ".terraform/virtualbox/gold")
	}
	
	return &ProviderConfig{
		Manager:       manager,
		MachineFolder: machineFolder,
		GoldFolder:    goldFolder,
	}, nil
}
