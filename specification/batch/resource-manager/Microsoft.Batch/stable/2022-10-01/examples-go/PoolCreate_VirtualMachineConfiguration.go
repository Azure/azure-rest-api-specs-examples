package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/PoolCreate_VirtualMachineConfiguration.json
func ExamplePoolClient_Create_createPoolFullVirtualMachineConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPoolClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "default-azurebatch-japaneast", "sampleacct", "testpool", armbatch.Pool{
		Properties: &armbatch.PoolProperties{
			DeploymentConfiguration: &armbatch.DeploymentConfiguration{
				VirtualMachineConfiguration: &armbatch.VirtualMachineConfiguration{
					DataDisks: []*armbatch.DataDisk{
						{
							Caching:            to.Ptr(armbatch.CachingTypeReadWrite),
							DiskSizeGB:         to.Ptr[int32](30),
							Lun:                to.Ptr[int32](0),
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypePremiumLRS),
						},
						{
							Caching:            to.Ptr(armbatch.CachingTypeNone),
							DiskSizeGB:         to.Ptr[int32](200),
							Lun:                to.Ptr[int32](1),
							StorageAccountType: to.Ptr(armbatch.StorageAccountTypeStandardLRS),
						}},
					DiskEncryptionConfiguration: &armbatch.DiskEncryptionConfiguration{
						Targets: []*armbatch.DiskEncryptionTarget{
							to.Ptr(armbatch.DiskEncryptionTargetOsDisk),
							to.Ptr(armbatch.DiskEncryptionTargetTemporaryDisk)},
					},
					ImageReference: &armbatch.ImageReference{
						Offer:     to.Ptr("WindowsServer"),
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						SKU:       to.Ptr("2016-Datacenter-SmallDisk"),
						Version:   to.Ptr("latest"),
					},
					LicenseType:    to.Ptr("Windows_Server"),
					NodeAgentSKUID: to.Ptr("batch.node.windows amd64"),
					NodePlacementConfiguration: &armbatch.NodePlacementConfiguration{
						Policy: to.Ptr(armbatch.NodePlacementPolicyTypeZonal),
					},
					OSDisk: &armbatch.OSDisk{
						EphemeralOSDiskSettings: &armbatch.DiffDiskSettings{
							Placement: to.Ptr("CacheDisk"),
						},
					},
					WindowsConfiguration: &armbatch.WindowsConfiguration{
						EnableAutomaticUpdates: to.Ptr(false),
					},
				},
			},
			NetworkConfiguration: &armbatch.NetworkConfiguration{
				EndpointConfiguration: &armbatch.PoolEndpointConfiguration{
					InboundNatPools: []*armbatch.InboundNatPool{
						{
							Name:                   to.Ptr("testnat"),
							BackendPort:            to.Ptr[int32](12001),
							FrontendPortRangeEnd:   to.Ptr[int32](15100),
							FrontendPortRangeStart: to.Ptr[int32](15000),
							NetworkSecurityGroupRules: []*armbatch.NetworkSecurityGroupRule{
								{
									Access:              to.Ptr(armbatch.NetworkSecurityGroupRuleAccessAllow),
									Priority:            to.Ptr[int32](150),
									SourceAddressPrefix: to.Ptr("192.100.12.45"),
									SourcePortRanges: []*string{
										to.Ptr("1"),
										to.Ptr("2")},
								},
								{
									Access:              to.Ptr(armbatch.NetworkSecurityGroupRuleAccessDeny),
									Priority:            to.Ptr[int32](3500),
									SourceAddressPrefix: to.Ptr("*"),
									SourcePortRanges: []*string{
										to.Ptr("*")},
								}},
							Protocol: to.Ptr(armbatch.InboundEndpointProtocolTCP),
						}},
				},
			},
			ScaleSettings: &armbatch.ScaleSettings{
				AutoScale: &armbatch.AutoScaleSettings{
					EvaluationInterval: to.Ptr("PT5M"),
					Formula:            to.Ptr("$TargetDedicatedNodes=1"),
				},
			},
			VMSize: to.Ptr("STANDARD_D4"),
		},
	}, &armbatch.PoolClientCreateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
