package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/GetDiagnosticForKeyVaultResource.json
func ExampleDiagnosticsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticsClient().Get(ctx, "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read", "VMNotWorkingInsight", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticResource = armselfhelp.DiagnosticResource{
	// 	Name: to.Ptr("Microsoft.Help/diagnostics"),
	// 	Type: to.Ptr("VMNotWorkingInsight"),
	// 	ID: to.Ptr("/subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read/providers/Microsoft.Help/diagnostics/VMNotWorkingInsight"),
	// 	Properties: &armselfhelp.DiagnosticResourceProperties{
	// 		AcceptedAt: to.Ptr("2023-03-10T03:04:40Z"),
	// 		Diagnostics: []*armselfhelp.Diagnostic{
	// 			{
	// 				Error: &armselfhelp.Error{
	// 					Code: to.Ptr("errorCode"),
	// 					Message: to.Ptr("errorMessage"),
	// 				},
	// 				Insights: []*armselfhelp.Insight{
	// 					{
	// 						ID: to.Ptr("InsightArticleId"),
	// 						ImportanceLevel: to.Ptr(armselfhelp.ImportanceLevelCritical),
	// 						Results: to.Ptr("Article Content"),
	// 						Title: to.Ptr("An example title for an Insight"),
	// 				}},
	// 				SolutionID: to.Ptr("sampleSolutionId"),
	// 				Status: to.Ptr(armselfhelp.StatusSucceeded),
	// 		}},
	// 		ProvisioningState: to.Ptr(armselfhelp.DiagnosticProvisioningStateSucceeded),
	// 	},
	// }
}
