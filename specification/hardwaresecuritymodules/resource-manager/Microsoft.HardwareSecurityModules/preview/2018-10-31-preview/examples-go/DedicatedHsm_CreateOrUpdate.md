Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhardwaresecuritymodules%2Farmhardwaresecuritymodules%2Fv0.1.0/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.

```go
package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
)

// x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2018-10-31-preview/examples/DedicatedHsm_CreateOrUpdate.json
func ExampleDedicatedHsmClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhardwaresecuritymodules.NewDedicatedHsmClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armhardwaresecuritymodules.DedicatedHsm{
			Resource: armhardwaresecuritymodules.Resource{
				Location: to.StringPtr("<location>"),
				SKU: &armhardwaresecuritymodules.SKU{
					Name: armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790.ToPtr(),
				},
				Tags: map[string]*string{
					"Dept":        to.StringPtr("hsm"),
					"Environment": to.StringPtr("dogfood"),
				},
			},
			Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
				NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
						{
							PrivateIPAddress: to.StringPtr("<private-ipaddress>"),
						}},
					Subnet: &armhardwaresecuritymodules.APIEntityReference{
						ID: to.StringPtr("<id>"),
					},
				},
				StampID: to.StringPtr("<stamp-id>"),
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
```
