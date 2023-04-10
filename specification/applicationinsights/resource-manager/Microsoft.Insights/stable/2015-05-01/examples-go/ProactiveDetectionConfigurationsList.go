package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ProactiveDetectionConfigurationsList.json
func ExampleProactiveDetectionConfigurationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProactiveDetectionConfigurationsClient().List(ctx, "my-resource-group", "my-component", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentProactiveDetectionConfigurationArray = []*armapplicationinsights.ComponentProactiveDetectionConfiguration{
	// 	{
	// 		CustomEmails: []*string{
	// 			to.Ptr("foo@microsoft.com"),
	// 			to.Ptr("foo2@microsoft.com")},
	// 			Enabled: to.Ptr(true),
	// 			Name: to.Ptr("slowpageloadtime"),
	// 			RuleDefinitions: &armapplicationinsights.ComponentProactiveDetectionConfigurationRuleDefinitions{
	// 				Description: to.Ptr("Smart Detection rules notify you of performance anomaly issues."),
	// 				DisplayName: to.Ptr("Slow page load time"),
	// 				HelpURL: to.Ptr("https://docs.microsoft.com/en-us/azure/application-insights/app-insights-proactive-performance-diagnostics"),
	// 				IsEnabledByDefault: to.Ptr(true),
	// 				IsHidden: to.Ptr(false),
	// 				IsInPreview: to.Ptr(false),
	// 				Name: to.Ptr("slowpageloadtime"),
	// 				SupportsEmailNotifications: to.Ptr(true),
	// 			},
	// 			SendEmailsToSubscriptionOwners: to.Ptr(true),
	// 		},
	// 		{
	// 			CustomEmails: []*string{
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Name: to.Ptr("slowpageloadtime"),
	// 			RuleDefinitions: &armapplicationinsights.ComponentProactiveDetectionConfigurationRuleDefinitions{
	// 				Description: to.Ptr("Smart Detection rules notify you of performance anomaly issues."),
	// 				DisplayName: to.Ptr("Slow server response time"),
	// 				HelpURL: to.Ptr("https://docs.microsoft.com/en-us/azure/application-insights/app-insights-proactive-performance-diagnostics"),
	// 				IsEnabledByDefault: to.Ptr(true),
	// 				IsHidden: to.Ptr(false),
	// 				IsInPreview: to.Ptr(false),
	// 				Name: to.Ptr("slowserverresponsetime"),
	// 				SupportsEmailNotifications: to.Ptr(true),
	// 			},
	// 			SendEmailsToSubscriptionOwners: to.Ptr(true),
	// 	}}
}
