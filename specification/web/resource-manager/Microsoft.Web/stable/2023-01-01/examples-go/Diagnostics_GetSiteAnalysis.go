package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/Diagnostics_GetSiteAnalysis.json
func ExampleDiagnosticsClient_GetSiteAnalysis_getAppAnalysis() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().GetSiteAnalysis(ctx, "Sample-WestUSResourceGroup", "SampleApp", "availability", "appanalysis", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AnalysisDefinition = armappservice.AnalysisDefinition{
	// 	Name: to.Ptr("appanalysis"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/Sample-WestUSResourceGroup/providers/Microsoft.Web/sites/SampleApp/slots/staging/diagnostics/availability/analyses/appanalysis"),
	// 	Properties: &armappservice.AnalysisDefinitionProperties{
	// 		Description: to.Ptr("Determine causes of availability loss as well as solutions for these problems"),
	// 	},
	// }
}
