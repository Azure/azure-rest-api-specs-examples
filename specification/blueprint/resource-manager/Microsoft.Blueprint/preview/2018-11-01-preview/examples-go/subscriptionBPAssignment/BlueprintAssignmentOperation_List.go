package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignmentOperation_List.json
func ExampleAssignmentOperationsClient_NewListPager_assignmentAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblueprint.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentOperationsClient().NewListPager("subscriptions/00000000-0000-0000-0000-000000000000", "assignSimpleBlueprint", nil)
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
		// page.AssignmentOperationList = armblueprint.AssignmentOperationList{
		// 	Value: []*armblueprint.AssignmentOperation{
		// 		{
		// 			Name: to.Ptr("fb5d4dcb-7ce2-4087-ba7a-459aa74e5e0f"),
		// 			Type: to.Ptr("microsoft.blueprint/blueprintAssignments/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.blueprint/blueprintAssignments/assignSimpleBlueprint/assignmentOperations/fb5d4dcb-7ce2-4087-ba7a-459aa74e5e0f"),
		// 			Properties: &armblueprint.AssignmentOperationProperties{
		// 				AssignmentState: to.Ptr("succeed"),
		// 				BlueprintVersion: to.Ptr("v20181101"),
		// 				Deployments: []*armblueprint.AssignmentDeploymentJob{
		// 					{
		// 						Action: to.Ptr("put"),
		// 						History: []*armblueprint.AssignmentDeploymentJobResult{
		// 							{
		// 								Error: &armblueprint.AzureResourceManagerError{
		// 									Code: to.Ptr("dummy"),
		// 									Message: to.Ptr("dummy"),
		// 								},
		// 						}},
		// 						JobState: to.Ptr("succeeded"),
		// 						Kind: to.Ptr("azureResource"),
		// 						RequestURI: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.deployments/deployments/48432786-2f1b-4925-8032-a5d57bcb5b6e"),
		// 						Result: &armblueprint.AssignmentDeploymentJobResult{
		// 							Resources: []*armblueprint.AssignmentJobCreatedResource{
		// 								{
		// 									Name: to.Ptr("foobar"),
		// 									Type: to.Ptr("foo/bar"),
		// 									ID: to.Ptr("blabla"),
		// 							}},
		// 						},
		// 				}},
		// 				TimeCreated: to.Ptr("2018-11-13T15:19:45-08:00"),
		// 				TimeFinished: to.Ptr("2018-11-13T15:26:02-08:00"),
		// 				TimeStarted: to.Ptr("2018-11-13T15:21:49-08:00"),
		// 			},
		// 	}},
		// }
	}
}
