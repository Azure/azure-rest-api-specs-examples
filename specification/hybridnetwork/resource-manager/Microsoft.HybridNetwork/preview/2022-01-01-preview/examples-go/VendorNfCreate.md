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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorNfCreate.json
func ExampleVendorNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewVendorNetworkFunctionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<location-name>",
		"<vendor-name>",
		"<service-key>",
		armhybridnetwork.VendorNetworkFunction{
			Properties: &armhybridnetwork.VendorNetworkFunctionPropertiesFormat{
				NetworkFunctionVendorConfigurations: []*armhybridnetwork.NetworkFunctionVendorConfiguration{
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
						OSProfile: &armhybridnetwork.OsProfile{
							AdminUsername: to.Ptr("<admin-username>"),
							CustomData:    to.Ptr("<custom-data>"),
							LinuxConfiguration: &armhybridnetwork.LinuxConfiguration{
								SSH: &armhybridnetwork.SSHConfiguration{
									PublicKeys: []*armhybridnetwork.SSHPublicKey{
										{
											Path:    to.Ptr("<path>"),
											KeyData: to.Ptr("<key-data>"),
										}},
								},
							},
						},
						RoleName:           to.Ptr("<role-name>"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUType:                 to.Ptr(armhybridnetwork.SKUTypeSDWAN),
				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateProvisioning),
			},
		},
		&armhybridnetwork.VendorNetworkFunctionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
