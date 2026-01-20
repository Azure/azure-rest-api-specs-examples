package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachineExtension_Get_MaximumSet_Gen.json
func ExampleVirtualMachineExtensionsClient_Get_virtualMachineExtensionGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineExtensionsClient().Get(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", &armcompute.VirtualMachineExtensionsClientGetOptions{Expand: to.Ptr("aaaaaa")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineExtension = armcompute.VirtualMachineExtension{
	// 	Name: to.Ptr("myVMExtension"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/myVMExtension"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key9183": to.Ptr("aa"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineExtensionProperties{
	// 		Type: to.Ptr("extType"),
	// 		AutoUpgradeMinorVersion: to.Ptr(true),
	// 		EnableAutomaticUpgrade: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("a"),
	// 		InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr("aaaaaaaaa"),
	// 			Statuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 					DisplayStatus: to.Ptr("aaaaaa"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Message: to.Ptr("a"),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 			}},
	// 			Substatuses: []*armcompute.InstanceViewStatus{
	// 				{
	// 					Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 					DisplayStatus: to.Ptr("aaaaaa"),
	// 					Level: to.Ptr(armcompute.StatusLevelTypesInfo),
	// 					Message: to.Ptr("a"),
	// 					Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
	// 			}},
	// 			TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		ProtectedSettings: map[string]any{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Publisher: to.Ptr("extPublisher"),
	// 		Settings: map[string]any{
	// 		},
	// 		SuppressFailures: to.Ptr(true),
	// 		TypeHandlerVersion: to.Ptr("1.2"),
	// 	},
	// }
}
