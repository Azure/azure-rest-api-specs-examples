package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/GetRunOutput.json
func ExampleVirtualMachineImageTemplatesClient_GetRunOutput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineImageTemplatesClient().GetRunOutput(ctx, "myResourceGroup", "myImageTemplate", "myManagedImageOutput", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RunOutput = armvirtualmachineimagebuilder.RunOutput{
	// 	Name: to.Ptr("myManagedImageOutput"),
	// 	Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/runOutputs"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate/runOutputs/myManagedImageOutput"),
	// 	Properties: &armvirtualmachineimagebuilder.RunOutputProperties{
	// 		ArtifactID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/output_managed_image"),
	// 		ProvisioningState: to.Ptr(armvirtualmachineimagebuilder.ProvisioningStateSucceeded),
	// 	},
	// }
}
