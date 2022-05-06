Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.4.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationUpdate.json
func ExampleProactiveDetectionConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewProactiveDetectionConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<configuration-id>",
		armapplicationinsights.ComponentProactiveDetectionConfiguration{
			CustomEmails: []*string{
				to.Ptr("foo@microsoft.com"),
				to.Ptr("foo2@microsoft.com")},
			Enabled: to.Ptr(true),
			Name:    to.Ptr("<name>"),
			RuleDefinitions: &armapplicationinsights.ComponentProactiveDetectionConfigurationRuleDefinitions{
				Description:                to.Ptr("<description>"),
				DisplayName:                to.Ptr("<display-name>"),
				HelpURL:                    to.Ptr("<help-url>"),
				IsEnabledByDefault:         to.Ptr(true),
				IsHidden:                   to.Ptr(false),
				IsInPreview:                to.Ptr(false),
				Name:                       to.Ptr("<name>"),
				SupportsEmailNotifications: to.Ptr(true),
			},
			SendEmailsToSubscriptionOwners: to.Ptr(true),
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
