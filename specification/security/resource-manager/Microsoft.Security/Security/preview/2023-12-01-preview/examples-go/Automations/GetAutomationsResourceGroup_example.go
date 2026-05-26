package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2023-12-01-preview/Automations/GetAutomationsResourceGroup_example.json
func ExampleAutomationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAutomationsClient().NewListByResourceGroupPager("exampleResourceGroup", nil)
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
		// page = armsecurity.AutomationsClientListByResourceGroupResponse{
		// 	AutomationList: armsecurity.AutomationList{
		// 		Value: []*armsecurity.Automation{
		// 			{
		// 				Name: to.Ptr("exampleAutomation"),
		// 				Type: to.Ptr("Microsoft.Security/automations"),
		// 				Etag: to.Ptr("etag value"),
		// 				ID: to.Ptr("/subscriptions/e4272367-5645-4c4e-9c67-3b74b59a6982/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
		// 				Location: to.Ptr("Central US"),
		// 				Properties: &armsecurity.AutomationProperties{
		// 					Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
		// 					Actions: []armsecurity.AutomationActionClassification{
		// 						&armsecurity.AutomationActionLogicApp{
		// 							ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
		// 							LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
		// 						},
		// 					},
		// 					IsEnabled: to.Ptr(true),
		// 					Scopes: []*armsecurity.AutomationScope{
		// 						{
		// 							Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
		// 							ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
		// 						},
		// 					},
		// 					Sources: []*armsecurity.AutomationSource{
		// 						{
		// 							EventSource: to.Ptr(armsecurity.EventSourceAssessments),
		// 							RuleSets: []*armsecurity.AutomationRuleSet{
		// 								{
		// 									Rules: []*armsecurity.AutomationTriggeringRule{
		// 										{
		// 											ExpectedValue: to.Ptr("customAssessment"),
		// 											Operator: to.Ptr(armsecurity.OperatorEquals),
		// 											PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
		// 											PropertyType: to.Ptr(armsecurity.PropertyTypeString),
		// 										},
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
