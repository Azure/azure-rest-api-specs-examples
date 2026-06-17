package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/GoalTemplates_List_MaximumSet_Gen.json
func ExampleGoalTemplatesClient_NewListPager_goalTemplatesListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGoalTemplatesClient().NewListPager("vmmacokmkuxzy", &armresiliencemanagement.GoalTemplatesClientListOptions{
		SkipToken: to.Ptr("xntbyoswztnmvitj"),
		Top:       to.Ptr[int32](69)})
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
		// 					RequireHighAvailability: to.Ptr(armresiliencemanagement.RequirementSelectedRequired),
		// 					RequireDisasterRecovery: to.Ptr(armresiliencemanagement.RequirementSelectedNotRequired),
		// 					RegionalRecoveryPointObjective: to.Ptr("PT15M"),
		// 					RegionalRecoveryTimeObjective: to.Ptr("PT30M"),
		// 					GoalType: to.Ptr(armresiliencemanagement.GoalTypeResiliency),
		// 					ProvisioningState: to.Ptr(armresiliencemanagement.ProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/providers/Microsoft.AzureResilienceManagement/goaltemplates/gt1"),
		// 				Name: to.Ptr("gt1"),
		// 				Type: to.Ptr("Microsoft.AzureResilienceManagement/goaltemplates"),
		// 				SystemData: &armresiliencemanagement.SystemData{
		// 					CreatedBy: to.Ptr("dvnfxbuyqhvivfjddjccdtlwajfht"),
		// 					CreatedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.796Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("lndhhaimomorael"),
		// 					LastModifiedByType: to.Ptr(armresiliencemanagement.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-06T15:03:42.797Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/akqqj"),
		// 	},
		// }
	}
}
