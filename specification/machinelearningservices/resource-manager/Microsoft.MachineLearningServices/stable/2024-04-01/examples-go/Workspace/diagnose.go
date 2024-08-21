package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Workspace/diagnose.json
func ExampleWorkspacesClient_BeginDiagnose() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginDiagnose(ctx, "workspace-1234", "testworkspace", &armmachinelearning.WorkspacesClientBeginDiagnoseOptions{Parameters: &armmachinelearning.DiagnoseWorkspaceParameters{
		Value: &armmachinelearning.DiagnoseRequestProperties{
			ApplicationInsights: map[string]any{},
			ContainerRegistry:   map[string]any{},
			DNSResolution:       map[string]any{},
			KeyVault:            map[string]any{},
			Nsg:                 map[string]any{},
			Others:              map[string]any{},
			ResourceLock:        map[string]any{},
			StorageAccount:      map[string]any{},
			Udr:                 map[string]any{},
		},
	},
	})
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
	// res.DiagnoseResponseResult = armmachinelearning.DiagnoseResponseResult{
	// 	Value: &armmachinelearning.DiagnoseResponseResultValue{
	// 		ApplicationInsightsResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		ContainerRegistryResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		DNSResolutionResults: []*armmachinelearning.DiagnoseResult{
	// 			{
	// 				Code: to.Ptr("CustomDNSInUse"),
	// 				Level: to.Ptr(armmachinelearning.DiagnoseResultLevelWarning),
	// 				Message: to.Ptr("We have detected an on-premise dns server is configured. Please make sure conditional forwarding is configured correctly according to doc https://foo"),
	// 		}},
	// 		KeyVaultResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		NetworkSecurityRuleResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		OtherResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		ResourceLockResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		StorageAccountResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 		UserDefinedRouteResults: []*armmachinelearning.DiagnoseResult{
	// 		},
	// 	},
	// }
}
