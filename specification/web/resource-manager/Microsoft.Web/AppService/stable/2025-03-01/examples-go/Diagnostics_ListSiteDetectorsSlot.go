package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/Diagnostics_ListSiteDetectorsSlot.json
func ExampleDiagnosticsClient_NewListSiteDetectorsSlotPager_listAppSlotDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiagnosticsClient().NewListSiteDetectorsSlotPager("Sample-WestUSResourceGroup", "SampleApp", "availability", "staging", nil)
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
		// page.DiagnosticDetectorCollection = armappservice.DiagnosticDetectorCollection{
		// 	Value: []*armappservice.DetectorDefinitionResource{
		// 		{
		// 			Name: to.Ptr("servicehealth"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/slots/staging/diagnostics/availability/detectors/servicehealth"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Service Health"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("siteswap"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/slots/staging/diagnostics/availability/detectors/siteswap"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Swap Operations"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](8),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sitecrashes"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/slots/staging/diagnostics/availability/detectors/sitecrashes"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Crash Events"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](9),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deployment"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/slots/staging/diagnostics/availability/detectors/deployment"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Deployments"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](7),
		// 			},
		// 	}},
		// }
	}
}
