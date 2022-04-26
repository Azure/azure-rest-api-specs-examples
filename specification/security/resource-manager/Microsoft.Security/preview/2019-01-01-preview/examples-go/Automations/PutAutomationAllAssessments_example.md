Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewAutomationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-name>",
		armsecurity.Automation{
			Location: to.Ptr("<location>"),
			Etag:     to.Ptr("<etag>"),
			Tags:     map[string]*string{},
			Properties: &armsecurity.AutomationProperties{
				Description: to.Ptr("<description>"),
				Actions: []armsecurity.AutomationActionClassification{
					&armsecurity.AutomationActionLogicApp{
						ActionType:         to.Ptr(armsecurity.ActionTypeLogicApp),
						LogicAppResourceID: to.Ptr("<logic-app-resource-id>"),
						URI:                to.Ptr("<uri>"),
					}},
				IsEnabled: to.Ptr(true),
				Scopes: []*armsecurity.AutomationScope{
					{
						Description: to.Ptr("<description>"),
						ScopePath:   to.Ptr("<scope-path>"),
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
		return
	}
	// TODO: use response item
	_ = res
}
```
