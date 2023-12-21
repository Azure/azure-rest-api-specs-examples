package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/ListRunOutputs.json
func ExampleVirtualMachineImageTemplatesClient_NewListRunOutputsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineImageTemplatesClient().NewListRunOutputsPager("myResourceGroup", "myImageTemplate", nil)
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
		// page.RunOutputCollection = armvirtualmachineimagebuilder.RunOutputCollection{
		// 	Value: []*armvirtualmachineimagebuilder.RunOutput{
		// 		{
		// 			Name: to.Ptr("myManagedImageOutput"),
		// 			Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/runOutputs"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate/runOutputs/myManagedImageOutput"),
		// 			Properties: &armvirtualmachineimagebuilder.RunOutputProperties{
		// 				ArtifactID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/output_managed_image"),
		// 				ProvisioningState: to.Ptr(armvirtualmachineimagebuilder.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mySharedImageOutput"),
		// 			Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/runOutputs"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate/runOutputs/mySharedImageOutput"),
		// 			Properties: &armvirtualmachineimagebuilder.RunOutputProperties{
		// 				ArtifactID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/galleries/Gallery1/images/SharedImageOutput/imageversions/1.2.3"),
		// 				ProvisioningState: to.Ptr(armvirtualmachineimagebuilder.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
