package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a0168458930c86636a76bcd7acfdc9c81291bfc/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FeatureCapabilitiesGet.json
func ExampleComponentFeatureCapabilitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewComponentFeatureCapabilitiesClient().Get(ctx, "my-resource-group", "my-component", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentFeatureCapabilities = armapplicationinsights.ComponentFeatureCapabilities{
	// 	AnalyticsIntegration: to.Ptr(true),
	// 	APIAccessLevel: to.Ptr("Premium"),
	// 	ApplicationMap: to.Ptr(true),
	// 	BurstThrottlePolicy: to.Ptr("B2"),
	// 	DailyCap: to.Ptr[float32](0.0323),
	// 	DailyCapResetTime: to.Ptr[float32](4),
	// 	LiveStreamMetrics: to.Ptr(true),
	// 	MultipleStepWebTest: to.Ptr(true),
	// 	OpenSchema: to.Ptr(true),
	// 	PowerBIIntegration: to.Ptr(true),
	// 	ProactiveDetection: to.Ptr(false),
	// 	SupportExportData: to.Ptr(true),
	// 	ThrottleRate: to.Ptr[float32](0),
	// 	TrackingType: to.Ptr("Basic"),
	// 	WorkItemIntegration: to.Ptr(true),
	// }
}
