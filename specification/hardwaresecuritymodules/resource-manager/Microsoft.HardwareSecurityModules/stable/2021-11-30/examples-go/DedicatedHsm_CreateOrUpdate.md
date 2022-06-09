```go
package armhardwaresecuritymodules_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/DedicatedHsm_CreateOrUpdate.json
func ExampleDedicatedHsmClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhardwaresecuritymodules.NewDedicatedHsmClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"hsm-group",
		"hsm1",
		armhardwaresecuritymodules.DedicatedHsm{
			Location: to.Ptr("westus"),
			SKU: &armhardwaresecuritymodules.SKU{
				Name: to.Ptr(armhardwaresecuritymodules.SKUNameSafeNetLunaNetworkHSMA790),
			},
			Tags: map[string]*string{
				"Dept":        to.Ptr("hsm"),
				"Environment": to.Ptr("dogfood"),
			},
			Properties: &armhardwaresecuritymodules.DedicatedHsmProperties{
				NetworkProfile: &armhardwaresecuritymodules.NetworkProfile{
					NetworkInterfaces: []*armhardwaresecuritymodules.NetworkInterface{
						{
							PrivateIPAddress: to.Ptr("1.0.0.1"),
						}},
					Subnet: &armhardwaresecuritymodules.APIEntityReference{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"),
					},
				},
				StampID: to.Ptr("stamp01"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhardwaresecuritymodules%2Farmhardwaresecuritymodules%2Fv1.0.0/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/README.md) on how to add the SDK to your project and authenticate.
