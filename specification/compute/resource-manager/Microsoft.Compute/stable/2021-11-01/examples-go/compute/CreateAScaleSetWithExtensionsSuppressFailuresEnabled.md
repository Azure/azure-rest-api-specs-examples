Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAScaleSetWithExtensionsSuppressFailuresEnabled.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSet{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.VirtualMachineScaleSetProperties{
				Overprovision: to.BoolPtr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					Mode: armcompute.UpgradeModeManual.ToPtr(),
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.StringPtr("<type>"),
									AutoUpgradeMinorVersion: to.BoolPtr(false),
									Publisher:               to.StringPtr("<publisher>"),
									Settings:                map[string]interface{}{},
									SuppressFailures:        to.BoolPtr(true),
									TypeHandlerVersion:      to.StringPtr("<type-handler-version>"),
								},
							}},
					},
					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.BoolPtr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.StringPtr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &armcompute.APIEntityReference{
													ID: to.StringPtr("<id>"),
												},
											},
										}},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.StringPtr("<admin-password>"),
						AdminUsername:      to.StringPtr("<admin-username>"),
						ComputerNamePrefix: to.StringPtr("<computer-name-prefix>"),
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
						ImageReference: &armcompute.ImageReference{
							Offer:     to.StringPtr("<offer>"),
							Publisher: to.StringPtr("<publisher>"),
							SKU:       to.StringPtr("<sku>"),
							Version:   to.StringPtr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
							Caching:      armcompute.CachingTypesReadWrite.ToPtr(),
							CreateOption: armcompute.DiskCreateOptionTypes("FromImage").ToPtr(),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
							},
						},
					},
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(3),
				Tier:     to.StringPtr("<tier>"),
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
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientCreateOrUpdateResult)
}
```
