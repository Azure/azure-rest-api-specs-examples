Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateAScaleSetWithExtensionsSuppressFailuresEnabled.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSet{
			Location: to.Ptr("<location>"),
			Properties: &armcompute.VirtualMachineScaleSetProperties{
				Overprovision: to.Ptr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					Mode: to.Ptr(armcompute.UpgradeModeManual),
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.Ptr(true),
							StorageURI: to.Ptr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.Ptr("<type>"),
									AutoUpgradeMinorVersion: to.Ptr(false),
									Publisher:               to.Ptr("<publisher>"),
									Settings:                map[string]interface{}{},
									SuppressFailures:        to.Ptr(true),
									TypeHandlerVersion:      to.Ptr("<type-handler-version>"),
								},
							}},
					},
					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
									EnableIPForwarding: to.Ptr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.Ptr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &armcompute.APIEntityReference{
													ID: to.Ptr("<id>"),
												},
											},
										}},
									Primary: to.Ptr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
						AdminPassword:      to.Ptr("<admin-password>"),
						AdminUsername:      to.Ptr("<admin-username>"),
						ComputerNamePrefix: to.Ptr("<computer-name-prefix>"),
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
						ImageReference: &armcompute.ImageReference{
							Offer:     to.Ptr("<offer>"),
							Publisher: to.Ptr("<publisher>"),
							SKU:       to.Ptr("<sku>"),
							Version:   to.Ptr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
							Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
						},
					},
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](3),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
