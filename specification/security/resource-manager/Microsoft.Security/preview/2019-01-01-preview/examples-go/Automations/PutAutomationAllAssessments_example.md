Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/PutAutomationAllAssessments_example.json
func ExampleAutomationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAutomationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-name>",
		armsecurity.Automation{
			Location: to.StringPtr("<location>"),
			Etag:     to.StringPtr("<etag>"),
			Tags:     map[string]*string{},
			Properties: &armsecurity.AutomationProperties{
				Description: to.StringPtr("<description>"),
				Actions: []armsecurity.AutomationActionClassification{
					&armsecurity.AutomationActionLogicApp{
						ActionType:         armsecurity.ActionType("LogicApp").ToPtr(),
						LogicAppResourceID: to.StringPtr("<logic-app-resource-id>"),
						URI:                to.StringPtr("<uri>"),
					}},
				IsEnabled: to.BoolPtr(true),
				Scopes: []*armsecurity.AutomationScope{
					{
						Description: to.StringPtr("<description>"),
						ScopePath:   to.StringPtr("<scope-path>"),
					}},
				Sources: []*armsecurity.AutomationSource{
					{
						EventSource: armsecurity.EventSource("Assessments").ToPtr(),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AutomationsClientCreateOrUpdateResult)
}
```
