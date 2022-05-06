Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv0.4.0/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewNetworkFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<network-function-name>",
		armhybridnetwork.NetworkFunction{
			Location: to.Ptr("<location>"),
			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
				Device: &armhybridnetwork.SubResource{
					ID: to.Ptr("<id>"),
				},
				ManagedApplicationParameters: map[string]interface{}{},
				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
					{
						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr("<gateway>"),
										IPAddress:          to.Ptr("<ipaddress>"),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr("<subnet>"),
									}},
								MacAddress:           to.Ptr("<mac-address>"),
								NetworkInterfaceName: to.Ptr("<network-interface-name>"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
							},
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr("<gateway>"),
										IPAddress:          to.Ptr("<ipaddress>"),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr("<subnet>"),
									}},
								MacAddress:           to.Ptr("<mac-address>"),
								NetworkInterfaceName: to.Ptr("<network-interface-name>"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeWan),
							}},
						RoleName:           to.Ptr("<role-name>"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUName:    to.Ptr("<skuname>"),
				SKUType:    to.Ptr(armhybridnetwork.SKUTypeSDWAN),
				VendorName: to.Ptr("<vendor-name>"),
			},
		},
		&armhybridnetwork.NetworkFunctionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
