package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_NewListPager_virtualMachineScaleSetExtensionListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetExtensionsClient().NewListPager("rgcompute", "aaaaaaaaaaaaaaaaaaaa", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcompute.VirtualMachineScaleSetExtensionsClientListResponse{
		// 	VirtualMachineScaleSetExtensionListResult: armcompute.VirtualMachineScaleSetExtensionListResult{
		// 		Value: []*armcompute.VirtualMachineScaleSetExtension{
		// 			{
		// 				Name: to.Ptr("{extension-name}"),
		// 				Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 					AutoUpgradeMinorVersion: to.Ptr(true),
		// 					Publisher: to.Ptr("{extension-Publisher}"),
		// 					Type: to.Ptr("{extension-Type}"),
		// 					TypeHandlerVersion: to.Ptr("{handler-version}"),
		// 					Settings: map[string]any{
		// 					},
		// 					ForceUpdateTag: to.Ptr("aaaaaaaaa"),
		// 					EnableAutomaticUpgrade: to.Ptr(true),
		// 					ProtectedSettings: map[string]any{
		// 					},
		// 					ProvisioningState: to.Ptr("aaa"),
		// 					ProvisionAfterExtensions: []*string{
		// 						to.Ptr("aa"),
		// 					},
		// 					SuppressFailures: to.Ptr(true),
		// 				},
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				ID: to.Ptr("aaaaaaaa"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aa"),
		// 	},
		// }
	}
}
