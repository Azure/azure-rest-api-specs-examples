package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/Diagnostics_GetSiteDetector.json
func ExampleDiagnosticsClient_GetSiteDetector_getAppDetector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().GetSiteDetector(ctx, "Sample-WestUSResourceGroup", "SampleApp", "availability", "sitecrashes", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DetectorDefinitionResource = armappservice.DetectorDefinitionResource{
	// 	Name: to.Ptr("sitecrashes"),
	// 	ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/sitecrashes"),
	// 	Properties: &armappservice.DetectorDefinition{
	// 		DisplayName: to.Ptr("Site Crash Events"),
	// 		IsEnabled: to.Ptr(true),
	// 		Rank: to.Ptr[float64](9),
	// 	},
	// }
}
