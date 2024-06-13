package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/actions/GetAllActionsByAlertRule.json
func ExampleActionsClient_NewListByAlertRulePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActionsClient().NewListByAlertRulePager("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
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
		// page.ActionsList = armsecurityinsights.ActionsList{
		// 	Value: []*armsecurityinsights.ActionResponse{
		// 		{
		// 			Name: to.Ptr("912bec42-cb66-4c03-ac63-1761b6898c3e"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/alertRules/actions"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/73e01a99-5cd7-4139-a149-9f2736ff2ab5/actions/912bec42-cb66-4c03-ac63-1761b6898c3e"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.ActionResponseProperties{
		// 				LogicAppResourceID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.Logic/workflows/MyAlerts"),
		// 				WorkflowID: to.Ptr("cd3765391efd48549fd7681ded1d48d7"),
		// 			},
		// 	}},
		// }
	}
}
