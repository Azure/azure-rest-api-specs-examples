Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.2.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationUpdate.json
func ExampleProactiveDetectionConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewProactiveDetectionConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<configuration-id>",
		armapplicationinsights.ComponentProactiveDetectionConfiguration{
			CustomEmails: []*string{
				to.StringPtr("foo@microsoft.com"),
				to.StringPtr("foo2@microsoft.com")},
			Enabled: to.BoolPtr(true),
			Name:    to.StringPtr("<name>"),
			RuleDefinitions: &armapplicationinsights.ComponentProactiveDetectionConfigurationRuleDefinitions{
				Description:                to.StringPtr("<description>"),
				DisplayName:                to.StringPtr("<display-name>"),
				HelpURL:                    to.StringPtr("<help-url>"),
				IsEnabledByDefault:         to.BoolPtr(true),
				IsHidden:                   to.BoolPtr(false),
				IsInPreview:                to.BoolPtr(false),
				Name:                       to.StringPtr("<name>"),
				SupportsEmailNotifications: to.BoolPtr(true),
			},
			SendEmailsToSubscriptionOwners: to.BoolPtr(true),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProactiveDetectionConfigurationsClientUpdateResult)
}
```
