package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/60679ee3db06e93eb73faa0587fed93ed843d6dc/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetExtensionsClient_NewListPager_virtualMachineScaleSetExtensionListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.VirtualMachineScaleSetExtensionListResult = armcompute.VirtualMachineScaleSetExtensionListResult{
		// 	Value: []*armcompute.VirtualMachineScaleSetExtension{
		// 		{
		// 			ID: to.Ptr("aaaaaaaa"),
		// 			Name: to.Ptr("{extension-name}"),
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
		// 				Type: to.Ptr("{extension-Type}"),
		// 				AutoUpgradeMinorVersion: to.Ptr(true),
		// 				EnableAutomaticUpgrade: to.Ptr(true),
		// 				ForceUpdateTag: to.Ptr("aaaaaaaaa"),
		// 				ProtectedSettings: map[string]any{
		// 				},
		// 				ProvisionAfterExtensions: []*string{
		// 					to.Ptr("aa")},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					Publisher: to.Ptr("{extension-Publisher}"),
		// 					Settings: map[string]any{
		// 					},
		// 					SuppressFailures: to.Ptr(true),
		// 					TypeHandlerVersion: to.Ptr("{handler-version}"),
		// 				},
		// 		}},
		// 	}
	}
}
