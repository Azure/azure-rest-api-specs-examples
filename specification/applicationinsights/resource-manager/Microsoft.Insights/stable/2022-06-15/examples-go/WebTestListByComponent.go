package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestListByComponent.json
func ExampleWebTestsClient_NewListByComponentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebTestsClient().NewListByComponentPager("my-component", "my-resource-group", nil)
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
		// page.WebTestListResult = armapplicationinsights.WebTestListResult{
		// 	Value: []*armapplicationinsights.WebTest{
		// 		{
		// 			Name: to.Ptr("my-webtest-my-component"),
		// 			Type: to.Ptr("Microsoft.Insights/webtests"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/webtests/my-webtest-my-component"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Tags: map[string]*string{
		// 				"hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component": to.Ptr("Resource"),
		// 			},
		// 			Kind: to.Ptr(armapplicationinsights.WebTestKindPing),
		// 			Properties: &armapplicationinsights.WebTestProperties{
		// 				Configuration: &armapplicationinsights.WebTestPropertiesConfiguration{
		// 					WebTest: to.Ptr("<WebTest Name=\"my-webtest\" Id=\"678ddf96-1ab8-44c8-9274-123456789abc\" Enabled=\"True\" CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"120\" WorkItemIds=\"\" xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\" CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\" Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\"><Items><Request Method=\"GET\" Guid=\"a4162485-9114-fcfc-e086-123456789abc\" Version=\"1.1\" Url=\"http://my-component.azurewebsites.net\" ThinkTime=\"0\" Timeout=\"120\" ParseDependentRequests=\"True\" FollowRedirects=\"True\" RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\" ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\" IgnoreHttpStatusCode=\"False\" /></Items></WebTest>"),
		// 				},
		// 				Description: to.Ptr(""),
		// 				Enabled: to.Ptr(false),
		// 				Frequency: to.Ptr[int32](900),
		// 				WebTestKind: to.Ptr(armapplicationinsights.WebTestKindPing),
		// 				Locations: []*armapplicationinsights.WebTestGeolocation{
		// 					{
		// 						Location: to.Ptr("apac-hk-hkn-azr"),
		// 				}},
		// 				WebTestName: to.Ptr("my-webtest"),
		// 				RetryEnabled: to.Ptr(true),
		// 				SyntheticMonitorID: to.Ptr("my-webtest-my-component"),
		// 				Timeout: to.Ptr[int32](120),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
