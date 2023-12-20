package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/CreateDiagnosticForKeyVaultResource.json
func ExampleDiagnosticsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiagnosticsClient().BeginCreate(ctx, "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read", "VMNotWorkingInsight", armselfhelp.DiagnosticResource{
		Properties: &armselfhelp.DiagnosticResourceProperties{
			Insights: []*armselfhelp.DiagnosticInvocation{
				{
					SolutionID: to.Ptr("SampleSolutionId1"),
				},
				{
					SolutionID: to.Ptr("SampleSolutionId2"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticResource = armselfhelp.DiagnosticResource{
	// 	Name: to.Ptr("Microsoft.Help/diagnostics"),
	// 	Type: to.Ptr("VMNotWorkingInsight"),
	// 	ID: to.Ptr("/subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read/providers/Microsoft.Help/diagnostics/VMNotWorkingInsight"),
	// 	Properties: &armselfhelp.DiagnosticResourceProperties{
	// 		AcceptedAt: to.Ptr("2023-11-05T18:13:15.8708055+00:00"),
	// 		ProvisioningState: to.Ptr(armselfhelp.DiagnosticProvisioningStateSucceeded),
	// 	},
	// }
}
