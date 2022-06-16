package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationAllAssessments_example.json
func ExampleAutomationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAutomationsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleAutomation",
		armsecurity.Automation{
			Location: to.Ptr("Central US"),
			Etag:     to.Ptr("etag value (must be supplied for update)"),
			Tags:     map[string]*string{},
			Properties: &armsecurity.AutomationProperties{
				Description: to.Ptr("An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment"),
				Actions: []armsecurity.AutomationActionClassification{
					&armsecurity.AutomationActionLogicApp{
						ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
						LogicAppResourceID: to.Ptr("/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1"),
						URI:                to.Ptr("https://exampleTriggerUri1.com"),
					}},
				IsEnabled: to.Ptr(true),
				Scopes: []*armsecurity.AutomationScope{
					{
						Description: to.Ptr("A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5"),
						ScopePath:   to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup"),
					}},
				Sources: []*armsecurity.AutomationSource{
					{
						EventSource: to.Ptr(armsecurity.EventSourceAssessments),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
