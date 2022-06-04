Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv1.0.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestCreate.json
func ExampleWebTestsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWebTestsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"my-resource-group",
		"my-webtest-my-component",
		armapplicationinsights.WebTest{
			Location: to.Ptr("South Central US"),
			Kind:     to.Ptr(armapplicationinsights.WebTestKindPing),
			Properties: &armapplicationinsights.WebTestProperties{
				Configuration: &armapplicationinsights.WebTestPropertiesConfiguration{
					WebTest: to.Ptr("<WebTest Name=\"my-webtest\" Id=\"678ddf96-1ab8-44c8-9274-123456789abc\" Enabled=\"True\" CssProjectStructure=\"\" CssIteration=\"\" Timeout=\"120\" WorkItemIds=\"\" xmlns=\"http://microsoft.com/schemas/VisualStudio/TeamTest/2010\" Description=\"\" CredentialUserName=\"\" CredentialPassword=\"\" PreAuthenticate=\"True\" Proxy=\"default\" StopOnError=\"False\" RecordedResultFile=\"\" ResultsLocale=\"\" ><Items><Request Method=\"GET\" Guid=\"a4162485-9114-fcfc-e086-123456789abc\" Version=\"1.1\" Url=\"http://my-component.azurewebsites.net\" ThinkTime=\"0\" Timeout=\"120\" ParseDependentRequests=\"True\" FollowRedirects=\"True\" RecordResult=\"True\" Cache=\"False\" ResponseTimeGoal=\"0\" Encoding=\"utf-8\" ExpectedHttpStatusCode=\"200\" ExpectedResponseUrl=\"\" ReportingName=\"\" IgnoreHttpStatusCode=\"False\" /></Items></WebTest>"),
				},
				Description: to.Ptr("Ping web test alert for mytestwebapp"),
				Enabled:     to.Ptr(true),
				Frequency:   to.Ptr[int32](900),
				WebTestKind: to.Ptr(armapplicationinsights.WebTestKindPing),
				Locations: []*armapplicationinsights.WebTestGeolocation{
					{
						Location: to.Ptr("us-fl-mia-edge"),
					}},
				WebTestName:        to.Ptr("my-webtest-my-component"),
				RetryEnabled:       to.Ptr(true),
				SyntheticMonitorID: to.Ptr("my-webtest-my-component"),
				Timeout:            to.Ptr[int32](120),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
