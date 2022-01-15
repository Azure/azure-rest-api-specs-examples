Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv0.2.0/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridnetwork.NewNetworkFunctionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-function-name>",
		armhybridnetwork.NetworkFunction{
			Location: to.StringPtr("<location>"),
			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
				Device: &armhybridnetwork.SubResource{
					ID: to.StringPtr("<id>"),
				},
				ManagedApplicationParameters: map[string]interface{}{},
				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
					{
						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.StringPtr("<gateway>"),
										IPAddress:          to.StringPtr("<ipaddress>"),
										IPAllocationMethod: armhybridnetwork.IPAllocationMethod("Dynamic").ToPtr(),
										IPVersion:          armhybridnetwork.IPVersion("IPv4").ToPtr(),
										Subnet:             to.StringPtr("<subnet>"),
									}},
								MacAddress:           to.StringPtr("<mac-address>"),
								NetworkInterfaceName: to.StringPtr("<network-interface-name>"),
								VMSwitchType:         armhybridnetwork.VMSwitchType("Management").ToPtr(),
							},
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.StringPtr("<gateway>"),
										IPAddress:          to.StringPtr("<ipaddress>"),
										IPAllocationMethod: armhybridnetwork.IPAllocationMethod("Dynamic").ToPtr(),
										IPVersion:          armhybridnetwork.IPVersion("IPv4").ToPtr(),
										Subnet:             to.StringPtr("<subnet>"),
									}},
								MacAddress:           to.StringPtr("<mac-address>"),
								NetworkInterfaceName: to.StringPtr("<network-interface-name>"),
								VMSwitchType:         armhybridnetwork.VMSwitchType("Wan").ToPtr(),
							}},
						RoleName:           to.StringPtr("<role-name>"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUName:    to.StringPtr("<skuname>"),
				SKUType:    armhybridnetwork.SKUType("SDWAN").ToPtr(),
				VendorName: to.StringPtr("<vendor-name>"),
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
	log.Printf("Response result: %#v\n", res.NetworkFunctionsClientCreateOrUpdateResult)
}
```
