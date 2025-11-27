package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/Diagnostics_ListSiteDetectors.json
func ExampleDiagnosticsClient_NewListSiteDetectorsSlotPager_listAppDetectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiagnosticsClient().NewListSiteDetectorsSlotPager("Sample-WestUSResourceGroup", "SampleApp", "availability", "Production", nil)
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
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/servicehealth"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Service Health"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("siteswap"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/siteswap"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Swap Operations"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](8),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sitecrashes"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/sitecrashes"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Crash Events"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](9),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deployment"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/deployment"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Deployments"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](7),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sitecpuanalysis"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/sitecpuanalysis"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("CPU Analysis"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sitememoryanalysis"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/sitememoryanalysis"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Physical Memory Analysis"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("committedmemoryusage"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/committedmemoryusage"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Committed Memory Usage"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](7),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pagefileoperations"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/pagefileoperations"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Page File Operations"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](3),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("siterestartuserinitiated"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/siterestartuserinitiated"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("User Initiated Site Restarts"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](14),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("siterestartsettingupdate"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/siterestartsettingupdate"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Config Update Site Restarts"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](12),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("frebanalysis"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/frebanalysis"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Freb Logs Analysis"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](6),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("workeravailability"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/workeravailability"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Worker Availability"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](11),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sitelatency"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/sitelatency"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Site Latency"),
		// 				IsEnabled: to.Ptr(false),
		// 				Rank: to.Ptr[float64](1005),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("threadcount"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/threadcount"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Thread Count"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](23),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("failedrequestsperuri"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/failedrequestsperuri"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("Failed Requests Per URI"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](998),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("autoheal"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/autoheal"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("AutoHeal"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](21),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("loganalyzer"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/loganalyzer"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("PHP Log Analyzer"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](26),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aspnetcore"),
		// 			ID: to.Ptr("/subscriptions/ef90e930-9d7f-4a60-8a99-748e0eea69de/resourceGroups/Build2015DemoRG/providers/Microsoft.Web/sites/BuggyBakery/diagnostics/availability/detectors/aspnetcore"),
		// 			Properties: &armappservice.DetectorDefinition{
		// 				DisplayName: to.Ptr("ASP.NET Core"),
		// 				IsEnabled: to.Ptr(true),
		// 				Rank: to.Ptr[float64](5),
		// 			},
		// 	}},
		// }
	}
}
