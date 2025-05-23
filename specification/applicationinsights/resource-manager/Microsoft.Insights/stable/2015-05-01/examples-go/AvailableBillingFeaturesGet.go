package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AvailableBillingFeaturesGet.json
func ExampleComponentAvailableFeaturesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentAvailableFeaturesClient().Get(ctx, "my-resource-group", "my-component", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentAvailableFeatures = armapplicationinsights.ComponentAvailableFeatures{
	// 	Result: []*armapplicationinsights.ComponentFeature{
	// 		{
	// 			Capabilities: []*armapplicationinsights.ComponentFeatureCapability{
	// 				{
	// 					Description: to.Ptr("Number of application hosts"),
	// 					Name: to.Ptr("hostnumber"),
	// 					Value: to.Ptr("Unlimited"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Included data"),
	// 					MeterID: to.Ptr("acf26b15-ee92-440d-9973-9a72d77641aa"),
	// 					MeterRateFrequency: to.Ptr("GB/month"),
	// 					Name: to.Ptr("includeddata"),
	// 					Value: to.Ptr("1"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Additional data"),
	// 					MeterID: to.Ptr("b90f8b65-6c3e-43fc-9149-bdfc73b6a5b9"),
	// 					MeterRateFrequency: to.Ptr("/GB"),
	// 					Name: to.Ptr("additionaldata"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Data retention"),
	// 					Name: to.Ptr("dataretention"),
	// 					Unit: to.Ptr("days"),
	// 					Value: to.Ptr("90"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Response time monitoring and diagnostics"),
	// 					Name: to.Ptr("responsetimemonitoring"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Failed requests monitoring and diagnostics"),
	// 					Name: to.Ptr("failedrequestsmonitoring"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Browser performance"),
	// 					Name: to.Ptr("browserperformance"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Usage analysis"),
	// 					Name: to.Ptr("usageanalysis"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Server monitoring"),
	// 					Name: to.Ptr("servermonitoring"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Alerting and notifications"),
	// 					Name: to.Ptr("alertingandnotifications"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Daily notification of failed request rate spikes"),
	// 					Name: to.Ptr("notificationfailedrequestrate"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Telemetry analyzer"),
	// 					Name: to.Ptr("telemetryanalyzer"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Search and Analytics"),
	// 					Name: to.Ptr("searchandanalytics"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Web tests (multi-step tests)"),
	// 					MeterID: to.Ptr("0aa0e0e9-3f58-4dcf-9bb0-9db7ae1d5954"),
	// 					MeterRateFrequency: to.Ptr("/test (per month)"),
	// 					Name: to.Ptr("webtests"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Live stream metrics"),
	// 					Name: to.Ptr("livestreammetrics"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Application map"),
	// 					Name: to.Ptr("applicationmap"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Daily notification for many key metrics"),
	// 					Name: to.Ptr("dailynotificationforkeymetrics"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Work item integration"),
	// 					Name: to.Ptr("workitemintegration"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("API access"),
	// 					Name: to.Ptr("apiaccess"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Power BI integration"),
	// 					Name: to.Ptr("powerbiintegration"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Bulk data import"),
	// 					Name: to.Ptr("bulkdataimport"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Automatic data evaluation"),
	// 					Name: to.Ptr("automaticdataevaluation"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Analytics integration with Azure dashboards"),
	// 					Name: to.Ptr("analyticsintegration"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Continuous export"),
	// 					MeterID: to.Ptr("90fa4d31-3ea2-4178-a894-ec4c76c712b2"),
	// 					MeterRateFrequency: to.Ptr("/GB"),
	// 					Name: to.Ptr("continuousexport"),
	// 					Value: to.Ptr("Enabled"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Default daily cap"),
	// 					Name: to.Ptr("defaultdailycap"),
	// 					Unit: to.Ptr("G"),
	// 					Value: to.Ptr("100"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Default maximum daily cap"),
	// 					Name: to.Ptr("defaultmaxdailycap"),
	// 					Unit: to.Ptr("G"),
	// 					Value: to.Ptr("1000"),
	// 			}},
	// 			FeatureName: to.Ptr("Basic"),
	// 			IsHidden: to.Ptr(true),
	// 			IsMainFeature: to.Ptr(true),
	// 			MeterID: to.Ptr("c9a05f12-4910-4527-a9ec-1db4e4dba60e"),
	// 			MeterRateFrequency: to.Ptr("/month"),
	// 			SupportedAddonFeatures: to.Ptr("Application Insights Enterprise"),
	// 			Title: to.Ptr("Application Insights Basic"),
	// 		},
	// 		{
	// 			Capabilities: []*armapplicationinsights.ComponentFeatureCapability{
	// 				{
	// 					Description: to.Ptr("Enterprise Included data"),
	// 					MeterID: to.Ptr("acf26b15-ee92-440d-9973-9a72d77641aa"),
	// 					MeterRateFrequency: to.Ptr("GB/month"),
	// 					Name: to.Ptr("enterpriseincludeddata"),
	// 					Value: to.Ptr("0.20"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Enterprise Additional data"),
	// 					MeterID: to.Ptr("3fedc88a-b68f-4936-bbf0-f290a254388c"),
	// 					MeterRateFrequency: to.Ptr("/GB"),
	// 					Name: to.Ptr("enterpriseadditionaldata"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Default daily cap"),
	// 					Name: to.Ptr("defaultdailycap"),
	// 					Unit: to.Ptr("G"),
	// 					Value: to.Ptr("100"),
	// 				},
	// 				{
	// 					Description: to.Ptr("Default maximum daily cap"),
	// 					Name: to.Ptr("defaultmaxdailycap"),
	// 					Unit: to.Ptr("G"),
	// 					Value: to.Ptr("1000"),
	// 			}},
	// 			FeatureName: to.Ptr("Application Insights Enterprise"),
	// 			IsHidden: to.Ptr(false),
	// 			IsMainFeature: to.Ptr(false),
	// 			MeterID: to.Ptr("222f32c5-a319-4787-b934-5fb95105b2c8"),
	// 			MeterRateFrequency: to.Ptr("/node/month"),
	// 			Title: to.Ptr("Enterprise"),
	// 	}},
	// }
}
