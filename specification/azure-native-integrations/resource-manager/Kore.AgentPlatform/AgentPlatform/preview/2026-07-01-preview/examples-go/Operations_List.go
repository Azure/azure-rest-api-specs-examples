package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager_operationsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armkoreaiartemis.OperationsClientListResponse{
		// 	OperationListResult: armkoreaiartemis.OperationListResult{
		// 		Value: []*armkoreaiartemis.Operation{
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations"),
		// 					Operation: to.Ptr("Get Organization"),
		// 					Description: to.Ptr("Read an Organization resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/write"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations"),
		// 					Operation: to.Ptr("Create or Update Organization"),
		// 					Description: to.Ptr("Create or update an Organization resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/delete"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations"),
		// 					Operation: to.Ptr("Delete Organization"),
		// 					Description: to.Ptr("Delete an Organization resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/createWorkspace/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations"),
		// 					Operation: to.Ptr("Create Workspace"),
		// 					Description: to.Ptr("Create a Workspace under an Organization."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/deleteWorkspace/action"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations"),
		// 					Operation: to.Ptr("Delete Workspace"),
		// 					Description: to.Ptr("Delete a Workspace under an Organization."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/organizations/workspaces/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("organizations/workspaces"),
		// 					Operation: to.Ptr("Get Workspace"),
		// 					Description: to.Ptr("Read a Workspace resource."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Kore.AgentPlatform/operations/read"),
		// 				IsDataAction: to.Ptr(false),
		// 				Display: &armkoreaiartemis.OperationDisplay{
		// 					Provider: to.Ptr("Kore.AgentPlatform"),
		// 					Resource: to.Ptr("operations"),
		// 					Operation: to.Ptr("List Operations"),
		// 					Description: to.Ptr("List the operations supported by the Kore.AgentPlatform resource provider."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
