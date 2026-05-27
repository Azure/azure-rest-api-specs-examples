package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2021-08-01-preview/Assignments/ListBySubscriptionAssignments_example.json
func ExampleAssignmentsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListBySubscriptionPager(nil)
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
		// page = armsecurity.AssignmentsClientListBySubscriptionResponse{
		// 	AssignmentList: armsecurity.AssignmentList{
		// 		Value: []*armsecurity.Assignment{
		// 			{
		// 				Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Type: to.Ptr("Microsoft.Security/assignments"),
		// 				Etag: to.Ptr("etag value"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myResourceGroup/providers/Microsoft.Security/assignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armsecurity.AssignmentProperties{
		// 					Description: to.Ptr("Set of policies monitored by Azure Security Center for cross cloud"),
		// 					AdditionalData: &armsecurity.AssignmentPropertiesAdditionalData{
		// 						ExemptionCategory: to.Ptr("waiver"),
		// 					},
		// 					AssignedStandard: &armsecurity.AssignedStandardItem{
		// 						ID: to.Ptr("/providers/Microsoft.Security/Standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 					},
		// 					DisplayName: to.Ptr("ASC Default"),
		// 					Effect: to.Ptr("Exempt"),
		// 					ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T19:50:47.083633Z"); return t}()),
		// 					Metadata: map[string]any{
		// 						"ticketId": 12345,
		// 					},
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/ResourceGroup/rg"),
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Type: to.Ptr("Microsoft.Security/assignments"),
		// 				Etag: to.Ptr("etag value"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myResourceGroup/providers/Microsoft.Security/assignments/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armsecurity.AssignmentProperties{
		// 					Description: to.Ptr("Set of policies monitored by Azure Security Center for cross cloud"),
		// 					AdditionalData: &armsecurity.AssignmentPropertiesAdditionalData{
		// 						ExemptionCategory: to.Ptr("waiver"),
		// 					},
		// 					AssignedStandard: &armsecurity.AssignedStandardItem{
		// 						ID: to.Ptr("/providers/Microsoft.Security/Standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 					},
		// 					DisplayName: to.Ptr("ASC Default"),
		// 					Effect: to.Ptr("Exempt"),
		// 					ExpiresOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-05-01T19:50:47.083633Z"); return t}()),
		// 					Metadata: map[string]any{
		// 						"ticketId": 12345,
		// 					},
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/ResourceGroup/rg"),
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
