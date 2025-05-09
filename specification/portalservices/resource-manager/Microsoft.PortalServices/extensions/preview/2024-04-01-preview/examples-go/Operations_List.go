package armportalservicescopilot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portalservicescopilot/armportalservicescopilot"
)

// Generated from example definition: 2024-04-01-preview/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armportalservicescopilot.NewClientFactory(cred, nil)
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
		// page = armportalservicescopilot.OperationsClientListResponse{
		// 	OperationListResult: armportalservicescopilot.OperationListResult{
		// 		Value: []*armportalservicescopilot.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.PortalServices/copilotSettings/read"),
		// 				Display: &armportalservicescopilot.OperationDisplay{
		// 					Provider: to.Ptr("Azure Portal"),
		// 					Resource: to.Ptr("CopilotSettings"),
		// 					Operation: to.Ptr("Read copilotsettings"),
		// 					Description: to.Ptr("Read copilotsettings"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.PortalServices/copilotSettings/write"),
		// 				Display: &armportalservicescopilot.OperationDisplay{
		// 					Provider: to.Ptr("Azure Portal"),
		// 					Resource: to.Ptr("CopilotSettings"),
		// 					Operation: to.Ptr("Update copilotsettings"),
		// 					Description: to.Ptr("Update copilotsettings"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.PortalServices/copilotSettings/delete"),
		// 				Display: &armportalservicescopilot.OperationDisplay{
		// 					Provider: to.Ptr("Azure Portal"),
		// 					Resource: to.Ptr("CopilotSettings"),
		// 					Operation: to.Ptr("Delete copilotsettings"),
		// 					Description: to.Ptr("Delete copilotsettings"),
		// 				},
		// 			},
		// 			{
		// 				IsDataAction: to.Ptr(true),
		// 				Name: to.Ptr("Microsoft.PortalServices/copilotSettings/conversations/action"),
		// 				Display: &armportalservicescopilot.OperationDisplay{
		// 					Provider: to.Ptr("Azure Portal"),
		// 					Resource: to.Ptr("CopilotSettings"),
		// 					Operation: to.Ptr("Conversations"),
		// 					Description: to.Ptr("Conversations"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
