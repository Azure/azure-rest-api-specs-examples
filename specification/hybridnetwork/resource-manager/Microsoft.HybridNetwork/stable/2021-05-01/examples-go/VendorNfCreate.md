Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv0.2.1/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorNfCreate.json
func ExampleVendorNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridnetwork.NewVendorNetworkFunctionsClient("<subscription-id>", cred, nil)
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
						OSProfile: &armhybridnetwork.OsProfile{
							AdminUsername: to.StringPtr("<admin-username>"),
							CustomData:    to.StringPtr("<custom-data>"),
							LinuxConfiguration: &armhybridnetwork.LinuxConfiguration{
								SSH: &armhybridnetwork.SSHConfiguration{
									PublicKeys: []*armhybridnetwork.SSHPublicKey{
										{
											Path:    to.StringPtr("<path>"),
											KeyData: to.StringPtr("<key-data>"),
										}},
								},
							},
						},
						RoleName:           to.StringPtr("<role-name>"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUType:                 armhybridnetwork.SKUType("SDWAN").ToPtr(),
				VendorProvisioningState: armhybridnetwork.VendorProvisioningState("Provisioning").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.VendorNetworkFunctionsClientCreateOrUpdateResult)
}
```
