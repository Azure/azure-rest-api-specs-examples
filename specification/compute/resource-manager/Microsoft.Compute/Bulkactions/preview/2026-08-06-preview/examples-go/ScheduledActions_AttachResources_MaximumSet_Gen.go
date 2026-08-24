package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/ScheduledActions_AttachResources_MaximumSet_Gen.json
func ExampleScheduledActionsClient_BeginAttachResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("CB26D7CB-3E27-465F-99C8-EAF7A4118245", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScheduledActionsClient().BeginAttachResources(ctx, "rgcompute", "myScheduledAction", armbulkactions.ResourceAttachRequest{
		Resources: []*armbulkactions.ScheduledActionResourceInput{
			{
				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
				NotificationSettings: []*armbulkactions.NotificationProperties{
					{
						Destination: to.Ptr("admin@contoso.com"),
						Type:        to.Ptr(armbulkactions.NotificationTypeEmail),
						Language:    to.Ptr(armbulkactions.LanguageEnUs),
						Disabled:    to.Ptr(true),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.ScheduledActionsClientAttachResourcesResponse{
	// 	ResourceOperationResponse: armbulkactions.ResourceOperationResponse{
	// 		TotalResources: to.Ptr[int32](11),
	// 		ResourcesStatuses: []*armbulkactions.ResourceStatus{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/1d04e8f1-ee04-4056-b0b2-718f5bb45b04/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachines/myVm"),
	// 				Status: to.Ptr(armbulkactions.ResourceOperationStatusSucceeded),
	// 				Error: &armbulkactions.Error{
	// 					Code: to.Ptr("InternalServerError"),
	// 					Message: to.Ptr("An internal error occurred."),
	// 					Target: to.Ptr("virtualMachines"),
	// 					Details: []*armbulkactions.Error{
	// 					},
	// 					Innererror: &armbulkactions.InnerError{
	// 						Code: to.Ptr("InnerErrorCode"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
