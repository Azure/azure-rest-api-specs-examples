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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestCreate.json
func ExampleWebTestsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWebTestsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<web-test-name>",
		armapplicationinsights.WebTest{
			Location: to.Ptr("<location>"),
			Kind:     to.Ptr(armapplicationinsights.WebTestKindPing),
			Properties: &armapplicationinsights.WebTestProperties{
				Configuration: &armapplicationinsights.WebTestPropertiesConfiguration{
					WebTest: to.Ptr("<web-test>"),
				},
				Description: to.Ptr("<description>"),
				Enabled:     to.Ptr(true),
				Frequency:   to.Ptr[int32](900),
				WebTestKind: to.Ptr(armapplicationinsights.WebTestKindPing),
				Locations: []*armapplicationinsights.WebTestGeolocation{
					{
						Location: to.Ptr("<location>"),
					}},
				WebTestName:        to.Ptr("<web-test-name>"),
				RetryEnabled:       to.Ptr(true),
				SyntheticMonitorID: to.Ptr("<synthetic-monitor-id>"),
				Timeout:            to.Ptr[int32](120),
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
