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

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestCreate.json
func ExampleWebTestsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWebTestsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<web-test-name>",
		armapplicationinsights.WebTest{
			Location: to.StringPtr("<location>"),
			Kind:     armapplicationinsights.WebTestKindPing.ToPtr(),
			Properties: &armapplicationinsights.WebTestProperties{
				Configuration: &armapplicationinsights.WebTestPropertiesConfiguration{
					WebTest: to.StringPtr("<web-test>"),
				},
				Description: to.StringPtr("<description>"),
				Enabled:     to.BoolPtr(true),
				Frequency:   to.Int32Ptr(900),
				WebTestKind: armapplicationinsights.WebTestKindPing.ToPtr(),
				Locations: []*armapplicationinsights.WebTestGeolocation{
					{
						Location: to.StringPtr("<location>"),
					}},
				WebTestName:        to.StringPtr("<web-test-name>"),
				RetryEnabled:       to.BoolPtr(true),
				SyntheticMonitorID: to.StringPtr("<synthetic-monitor-id>"),
				Timeout:            to.Int32Ptr(120),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebTestsClientCreateOrUpdateResult)
}
```
