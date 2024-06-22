package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b22c642b361e6d6e7d72a2347a09b0bcf6075d70/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/GetTrigger.json
func ExampleTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().Get(ctx, "myResourceGroup", "myImageTemplate", "source", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Trigger = armvirtualmachineimagebuilder.Trigger{
	// 	Name: to.Ptr("source"),
	// 	Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/triggers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate/triggers/source"),
	// 	Properties: &armvirtualmachineimagebuilder.SourceImageTriggerProperties{
	// 		Kind: to.Ptr("SourceImage"),
	// 		ProvisioningState: to.Ptr(armvirtualmachineimagebuilder.ProvisioningStateSucceeded),
	// 		Status: &armvirtualmachineimagebuilder.TriggerStatus{
	// 			Code: to.Ptr("Healthy"),
	// 			Message: to.Ptr(""),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-21T17:32:28.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
