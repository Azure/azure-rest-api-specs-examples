```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationUpdate.json
func ExampleProactiveDetectionConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewProactiveDetectionConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"my-component",
		"slowpageloadtime",
		armapplicationinsights.ComponentProactiveDetectionConfiguration{
			CustomEmails: []*string{
				to.Ptr("foo@microsoft.com"),
				to.Ptr("foo2@microsoft.com")},
			Enabled: to.Ptr(true),
			Name:    to.Ptr("slowpageloadtime"),
			RuleDefinitions: &armapplicationinsights.ComponentProactiveDetectionConfigurationRuleDefinitions{
				Description:                to.Ptr("Smart Detection rules notify you of performance anomaly issues."),
				DisplayName:                to.Ptr("Slow page load time"),
				HelpURL:                    to.Ptr("https://docs.microsoft.com/en-us/azure/application-insights/app-insights-proactive-performance-diagnostics"),
				IsEnabledByDefault:         to.Ptr(true),
				IsHidden:                   to.Ptr(false),
				IsInPreview:                to.Ptr(false),
				Name:                       to.Ptr("slowpageloadtime"),
				SupportsEmailNotifications: to.Ptr(true),
			},
			SendEmailsToSubscriptionOwners: to.Ptr(true),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv2.0.0-beta.1/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.
