package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-08-31-preview/GoalTemplates_List_MinimumSet_Gen.json
func ExampleGoalTemplatesClient_NewListPager_goalTemplatesListMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoalTemplatesClient().NewListPager("sg1", nil)
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
		// page = armresiliencemanagement.GoalTemplatesClientListResponse{
		// 	GoalTemplateListResult: armresiliencemanagement.GoalTemplateListResult{
		// 		Value: []*armresiliencemanagement.GoalTemplate{
		// 			{
		// 				Properties: &armresiliencemanagement.GoalTemplateProperties{
		// 					GoalType: to.Ptr(armresiliencemanagement.GoalTypeResiliency),
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goaltemplates/gt1"),
		// 				Name: to.Ptr("gt1"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/goaltemplates"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("admin@contoso.com"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 796000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2025, time.February, 6, 15, 3, 42, 797000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
