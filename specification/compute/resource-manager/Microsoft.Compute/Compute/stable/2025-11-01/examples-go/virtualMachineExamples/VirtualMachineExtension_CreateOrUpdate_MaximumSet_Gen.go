package armcompute_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachineExtension_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionsClient_BeginCreateOrUpdate_virtualMachineExtensionCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineExtensionsClient().BeginCreateOrUpdate(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa", armcompute.VirtualMachineExtension{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineExtensionProperties{
			AutoUpgradeMinorVersion: to.Ptr(true),
			Publisher:               to.Ptr("extPublisher"),
			Type:                    to.Ptr("extType"),
			TypeHandlerVersion:      to.Ptr("1.2"),
			SuppressFailures:        to.Ptr(true),
			Settings:                map[string]any{},
			ForceUpdateTag:          to.Ptr("a"),
			EnableAutomaticUpgrade:  to.Ptr(true),
			ProtectedSettings:       map[string]any{},
			InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
				Name:               to.Ptr("aaaaaaaaaaaaaaaaa"),
				Type:               to.Ptr("aaaaaaaaa"),
				TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
				Substatuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						DisplayStatus: to.Ptr("aaaaaa"),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					},
				},
				Statuses: []*armcompute.InstanceViewStatus{
					{
						Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
						Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
						DisplayStatus: to.Ptr("aaaaaa"),
						Message:       to.Ptr("a"),
						Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
					},
				},
			},
		},
		Tags: map[string]*string{
			"key9183": to.Ptr("aa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineExtensionsClientCreateOrUpdateResponse{
	// 	VirtualMachineExtension: &armcompute.VirtualMachineExtension{
	// 		Name: to.Ptr("myVMExtension"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/myVMExtension"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcompute.VirtualMachineExtensionProperties{
	// 			AutoUpgradeMinorVersion: to.Ptr(true),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 			Publisher: to.Ptr("extPublisher"),
	// 			Type: to.Ptr("extType"),
	// 			TypeHandlerVersion: to.Ptr("1.2"),
	// 			SuppressFailures: to.Ptr(true),
	// 			Settings: map[string]any{
	// 			},
	// 			ForceUpdateTag: to.Ptr("a"),
	// 			EnableAutomaticUpgrade: to.Ptr(true),
	// 			ProtectedSettings: map[string]any{
	// 			},
	// 			InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
	// 				Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 				Type: to.Ptr("aaaaaaaaa"),
	// 				TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				Substatuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("aaaaaa"),
	// 						Message: to.Ptr("a"),
	// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 					},
	// 				},
	// 				Statuses: []*armcompute.InstanceViewStatus{
	// 					{
	// 						Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 						Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 						DisplayStatus: to.Ptr("aaaaaa"),
	// 						Message: to.Ptr("a"),
	// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key9183": to.Ptr("aa"),
	// 		},
	// 	},
	// }
}
