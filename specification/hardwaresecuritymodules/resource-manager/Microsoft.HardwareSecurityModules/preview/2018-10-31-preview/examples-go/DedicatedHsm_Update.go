package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
)

// x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2018-10-31-preview/examples/DedicatedHsm_Update.json
func ExampleDedicatedHsmClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhardwaresecuritymodules.NewDedicatedHsmClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armhardwaresecuritymodules.DedicatedHsmPatchParameters{
			Tags: map[string]*string{
				"Dept":        to.StringPtr("hsm"),
				"Environment": to.StringPtr("dogfood"),
				"Slice":       to.StringPtr("A"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DedicatedHsm.ID: %s\n", *res.ID)
}
