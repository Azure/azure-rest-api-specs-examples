package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/Diagnostics_ListSiteAnalysesSlot.json
func ExampleDiagnosticsClient_NewListSiteAnalysesPager_listAppSlotAnalyses() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiagnosticsClient().NewListSiteAnalysesPager("Sample-WestUSResourceGroup", "SampleApp", "availability", nil)
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
		// page.DiagnosticAnalysisCollection = armappservice.DiagnosticAnalysisCollection{
		// 	Value: []*armappservice.AnalysisDefinition{
		// 		{
		// 			Name: to.Ptr("appanalysis"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/appanalysis"),
		// 			Properties: &armappservice.AnalysisDefinitionProperties{
		// 				Description: to.Ptr("Determine causes of availability loss as well as solutions for these problems"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("apprestartanalysis"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/apprestartanalyses"),
		// 			Properties: &armappservice.AnalysisDefinitionProperties{
		// 				Description: to.Ptr("Find the reasons that your app restarted"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("memoryanalysis"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/memoryanalysis"),
		// 			Properties: &armappservice.AnalysisDefinitionProperties{
		// 				Description: to.Ptr("Detect issues with memory as well as suggest ways to troubleshoot memory problems"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("tcpconnectionsanalysis"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/tcpconnectionsanalysis"),
		// 			Properties: &armappservice.AnalysisDefinitionProperties{
		// 				Description: to.Ptr("Analyze port usage and find out if a high number of connections is causing problems for your web app"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("perfanalysis"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/perfanalysis"),
		// 			Properties: &armappservice.AnalysisDefinitionProperties{
		// 				Description: to.Ptr("Determine causes of performance degredation as well as solutions for these problems"),
		// 			},
		// 	}},
		// }
	}
}
