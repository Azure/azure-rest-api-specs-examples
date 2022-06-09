```go
package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorNfCreate.json
func ExampleVendorNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridnetwork.NewVendorNetworkFunctionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"eastus",
		"testVendor",
		"testServiceKey",
		armhybridnetwork.VendorNetworkFunction{
			Properties: &armhybridnetwork.VendorNetworkFunctionPropertiesFormat{
				NetworkFunctionVendorConfigurations: []*armhybridnetwork.NetworkFunctionVendorConfiguration{
					{
						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr(""),
										IPAddress:          to.Ptr(""),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr(""),
									}},
								MacAddress:           to.Ptr(""),
								NetworkInterfaceName: to.Ptr("nic1"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
							},
							{
								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
									{
										Gateway:            to.Ptr(""),
										IPAddress:          to.Ptr(""),
										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
										IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
										Subnet:             to.Ptr(""),
									}},
								MacAddress:           to.Ptr("DC-97-F8-79-16-7D"),
								NetworkInterfaceName: to.Ptr("nic2"),
								VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeWan),
							}},
						OSProfile: &armhybridnetwork.OsProfile{
							AdminUsername: to.Ptr("dummyuser"),
							CustomData:    to.Ptr("base-64 encoded string of custom data"),
							LinuxConfiguration: &armhybridnetwork.LinuxConfiguration{
								SSH: &armhybridnetwork.SSHConfiguration{
									PublicKeys: []*armhybridnetwork.SSHPublicKey{
										{
											Path:    to.Ptr("home/user/.ssh/authorized_keys"),
											KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAgEAwrr66r8n6B8Y0zMF3dOpXEapIQD9DiYQ6D6/zwor9o39jSkHNiMMER/GETBbzP83LOcekm02aRjo55ArO7gPPVvCXbrirJu9pkm4AC4BBre5xSLS= user@constoso-DSH"),
										}},
								},
							},
						},
						RoleName:           to.Ptr("testRole"),
						UserDataParameters: map[string]interface{}{},
					}},
				SKUType:                 to.Ptr(armhybridnetwork.SKUTypeSDWAN),
				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateProvisioning),
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridnetwork%2Farmhybridnetwork%2Fv1.0.0/sdk/resourcemanager/hybridnetwork/armhybridnetwork/README.md) on how to add the SDK to your project and authenticate.
