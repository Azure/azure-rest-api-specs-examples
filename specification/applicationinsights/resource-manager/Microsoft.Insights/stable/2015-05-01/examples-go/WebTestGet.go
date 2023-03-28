package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestGet.json
func ExampleWebTestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebTestsClient().Get(ctx, "my-resource-group", "my-webtest-01-mywebservice", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebTest = armapplicationinsights.WebTest{
	// 	Name: to.Ptr("my-webtest-01-mywebservice"),
	// 	Type: to.Ptr("Microsoft.Insights/webtests"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/my-test-resources/providers/Microsoft.Insights/webtests/my-webtest-01-mywebservice"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Tags: map[string]*string{
	// 		"hidden-link:/subscriptions/subid/resourceGroups/my-test-resources/providers/Microsoft.Insights/components/mytester": to.Ptr("Resource"),
	// 		"hidden-link:/subscriptions/subid/resourceGroups/my-test-resources/providers/Microsoft.Web/sites/mytester": to.Ptr("Resource"),
	// 	},
	// 	Kind: to.Ptr(armapplicationinsights.WebTestKindPing),
	// 	Properties: &armapplicationinsights.WebTestProperties{
	// 		Configuration: &armapplicationinsights.WebTestPropertiesConfiguration{
	// 			WebTest: to.Ptr("<WebTest Name=\"mytest-webtest-01\" Id=\"0317d26b-8672-4370-bd6b-123456789abc\" Enabled=\"True\" CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"30\" WorkItemIds=\"\" xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\" CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\" Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\"><Items><Request Method=\"GET\" Guid=\"a55ce143-4f1e-a7e6-b69e-123456789abc\" Version=\"1.1\" Url=\"http://mytester.azurewebsites.net\" ThinkTime=\"0\" Timeout=\"30\" ParseDependentRequests=\"False\" FollowRedirects=\"True\" RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\" ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\" IgnoreHttpStatusCode=\"False\" /></Items></WebTest>"),
	// 		},
	// 		Description: to.Ptr(""),
	// 		Enabled: to.Ptr(false),
	// 		Frequency: to.Ptr[int32](900),
	// 		WebTestKind: to.Ptr(armapplicationinsights.WebTestKindPing),
	// 		Locations: []*armapplicationinsights.WebTestGeolocation{
	// 			{
	// 				Location: to.Ptr("us-fl-mia-edge"),
	// 			},
	// 			{
	// 				Location: to.Ptr("apac-hk-hkn-azr"),
	// 		}},
	// 		WebTestName: to.Ptr("mytest-webtest-01"),
	// 		RetryEnabled: to.Ptr(true),
	// 		SyntheticMonitorID: to.Ptr("my-webtest-01-mywebservice"),
	// 		Timeout: to.Ptr[int32](30),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
