package armvirtualmachineimagebuilder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/CreateSourceImageTrigger.json
func ExampleTriggersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvirtualmachineimagebuilder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myImageTemplate", "source", armvirtualmachineimagebuilder.Trigger{
		Properties: &armvirtualmachineimagebuilder.SourceImageTriggerProperties{
			Kind: to.Ptr("SourceImage"),
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
	// res.Trigger = armvirtualmachineimagebuilder.Trigger{
	// 	Name: to.Ptr("source"),
	// 	Type: to.Ptr("Microsoft.VirtualMachineImages/imageTemplates/triggers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VirtualMachineImages/imageTemplates/myImageTemplate/triggers/source"),
	// 	Properties: &armvirtualmachineimagebuilder.SourceImageTriggerProperties{
	// 		Kind: to.Ptr("SourceImage"),
	// 		ProvisioningState: to.Ptr(armvirtualmachineimagebuilder.ProvisioningStateSucceeded),
	// 		Status: &armvirtualmachineimagebuilder.TriggerStatus{
	// 			Code: to.Ptr("Healthy"),
	// 			Message: to.Ptr(""),
	// 			Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-21T17:32:28.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
