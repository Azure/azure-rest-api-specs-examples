package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestCreateStandard.json
func ExampleWebTestsClient_CreateOrUpdate_webTestCreateStandard() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebTestsClient().CreateOrUpdate(ctx, "my-resource-group", "my-webtest-my-component", armapplicationinsights.WebTest{
		Location: to.Ptr("South Central US"),
		Properties: &armapplicationinsights.WebTestProperties{
			Description: to.Ptr("Ping web test alert for mytestwebapp"),
			Enabled:     to.Ptr(true),
			Frequency:   to.Ptr[int32](900),
			WebTestKind: to.Ptr(armapplicationinsights.WebTestKindStandard),
			Locations: []*armapplicationinsights.WebTestGeolocation{
				{
					Location: to.Ptr("us-fl-mia-edge"),
				}},
			WebTestName: to.Ptr("my-webtest-my-component"),
			Request: &armapplicationinsights.WebTestPropertiesRequest{
				Headers: []*armapplicationinsights.HeaderField{
					{
						HeaderFieldName:  to.Ptr("Content-Language"),
						HeaderFieldValue: to.Ptr("de-DE"),
					},
					{
						HeaderFieldName:  to.Ptr("Accept-Language"),
						HeaderFieldValue: to.Ptr("de-DE"),
					}},
				HTTPVerb:    to.Ptr("POST"),
				RequestBody: to.Ptr("SGVsbG8gd29ybGQ="),
				RequestURL:  to.Ptr("https://bing.com"),
			},
			RetryEnabled:       to.Ptr(true),
			SyntheticMonitorID: to.Ptr("my-webtest-my-component"),
			Timeout:            to.Ptr[int32](120),
			ValidationRules: &armapplicationinsights.WebTestPropertiesValidationRules{
				SSLCertRemainingLifetimeCheck: to.Ptr[int32](100),
				SSLCheck:                      to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WebTest = armapplicationinsights.WebTest{
	// 	Name: to.Ptr("my-webtest-my-component"),
	// 	Type: to.Ptr("Microsoft.Insights/webtests"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/webtests/my-webtest-my-component"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Tags: map[string]*string{
	// 		"hidden-link:/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component": to.Ptr("Resource"),
	// 	},
	// 	Properties: &armapplicationinsights.WebTestProperties{
	// 		Description: to.Ptr("Ping web test alert for mytestwebapp"),
	// 		Enabled: to.Ptr(true),
	// 		Frequency: to.Ptr[int32](900),
	// 		WebTestKind: to.Ptr(armapplicationinsights.WebTestKindStandard),
	// 		Locations: []*armapplicationinsights.WebTestGeolocation{
	// 			{
	// 				Location: to.Ptr("us-fl-mia-edge"),
	// 		}},
	// 		WebTestName: to.Ptr("my-webtest-my-component"),
	// 		Request: &armapplicationinsights.WebTestPropertiesRequest{
	// 			Headers: []*armapplicationinsights.HeaderField{
	// 				{
	// 					HeaderFieldName: to.Ptr("Content-Language"),
	// 					HeaderFieldValue: to.Ptr("de-DE"),
	// 				},
	// 				{
	// 					HeaderFieldName: to.Ptr("Accept-Language"),
	// 					HeaderFieldValue: to.Ptr("de-DE"),
	// 			}},
	// 			HTTPVerb: to.Ptr("POST"),
	// 			RequestBody: to.Ptr("SGVsbG8gd29ybGQ="),
	// 			RequestURL: to.Ptr("https://bing.com"),
	// 		},
	// 		RetryEnabled: to.Ptr(true),
	// 		SyntheticMonitorID: to.Ptr("my-webtest-my-component"),
	// 		Timeout: to.Ptr[int32](120),
	// 		ValidationRules: &armapplicationinsights.WebTestPropertiesValidationRules{
	// 			SSLCertRemainingLifetimeCheck: to.Ptr[int32](100),
	// 			SSLCheck: to.Ptr(true),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
