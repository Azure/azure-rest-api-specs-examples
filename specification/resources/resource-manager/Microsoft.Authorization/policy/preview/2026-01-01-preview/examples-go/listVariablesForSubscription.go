package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listVariablesForSubscription.json
func ExampleVariablesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVariablesClient().NewListPager(nil)
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
		// page = armpolicy.VariablesClientListResponse{
		// 	VariableListResult: armpolicy.VariableListResult{
		// 		Value: []*armpolicy.Variable{
		// 			{
		// 				Properties: &armpolicy.VariableProperties{
		// 					Columns: []*armpolicy.VariableColumn{
		// 						{
		// 							ColumnName: to.Ptr("TestColumn"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T01:01:01.1075056Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/variables/DemoTestVariable"),
		// 				Type: to.Ptr("Microsoft.Authorization/variables"),
		// 				Name: to.Ptr("DemoTestVariable"),
		// 			},
		// 			{
		// 				Properties: &armpolicy.VariableProperties{
		// 					Columns: []*armpolicy.VariableColumn{
		// 						{
		// 							ColumnName: to.Ptr("NetworkResourceName"),
		// 						},
		// 					},
		// 				},
		// 				SystemData: &armpolicy.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-07-01T01:01:01.1075056Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armpolicy.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T01:01:01.1075056Z"); return t}()),
		// 				},
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/variables/NetworkingVariable"),
		// 				Type: to.Ptr("Microsoft.Authorization/variables"),
		// 				Name: to.Ptr("NetworkingVariable"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
