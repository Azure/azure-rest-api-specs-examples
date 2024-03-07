package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutDisableAutomation_example.json
func ExampleAutomationsClient_CreateOrUpdate_disableOrEnableASecurityAutomation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutomationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleAutomation", armsecurity.Automation{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.AutomationProperties{
			Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
			Actions: []armsecurity.AutomationActionClassification{
				&armsecurity.AutomationActionLogicApp{
					ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
					LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
					URI:                to.Ptr("https://exampleTriggerUri1.com"),
				}},
			IsEnabled: to.Ptr(false),
			Scopes: []*armsecurity.AutomationScope{
				{
					Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
					ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
				}},
			Sources: []*armsecurity.AutomationSource{
				{
					EventSource: to.Ptr(armsecurity.EventSourceAssessments),
					RuleSets: []*armsecurity.AutomationRuleSet{
						{
							Rules: []*armsecurity.AutomationTriggeringRule{
								{
									ExpectedValue: to.Ptr("customAssessment"),
									Operator:      to.Ptr(armsecurity.OperatorEquals),
									PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
									PropertyType:  to.Ptr(armsecurity.PropertyTypeString),
								}},
						}},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Automation = armsecurity.Automation{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("new etag value"),
	// 	Name: to.Ptr("exampleAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/automations"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/automations/exampleAutomation"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.AutomationProperties{
	// 		Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment"),
	// 		Actions: []armsecurity.AutomationActionClassification{
	// 			&armsecurity.AutomationActionLogicApp{
	// 				ActionType: to.Ptr(armsecurity.ActionTypeLogicApp),
	// 				LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
	// 		}},
	// 		IsEnabled: to.Ptr(false),
	// 		Scopes: []*armsecurity.AutomationScope{
	// 			{
	// 				Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
	// 				ScopePath: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
	// 		}},
	// 		Sources: []*armsecurity.AutomationSource{
	// 			{
	// 				EventSource: to.Ptr(armsecurity.EventSourceAssessments),
	// 				RuleSets: []*armsecurity.AutomationRuleSet{
	// 					{
	// 						Rules: []*armsecurity.AutomationTriggeringRule{
	// 							{
	// 								ExpectedValue: to.Ptr("customAssessment"),
	// 								Operator: to.Ptr(armsecurity.OperatorEquals),
	// 								PropertyJPath: to.Ptr("$.Entity.AssessmentType"),
	// 								PropertyType: to.Ptr(armsecurity.PropertyTypeString),
	// 						}},
	// 				}},
	// 		}},
	// 	},
	// }
}
