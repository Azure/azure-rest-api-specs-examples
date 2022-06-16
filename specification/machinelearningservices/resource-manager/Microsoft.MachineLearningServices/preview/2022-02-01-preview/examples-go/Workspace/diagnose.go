package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Workspace/diagnose.json
func ExampleWorkspacesClient_BeginDiagnose() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDiagnose(ctx,
		"workspace-1234",
		"testworkspace",
		&armmachinelearning.WorkspacesClientBeginDiagnoseOptions{Parameters: &armmachinelearning.DiagnoseWorkspaceParameters{
			Value: &armmachinelearning.DiagnoseRequestProperties{
				ApplicationInsights: map[string]interface{}{},
				ContainerRegistry:   map[string]interface{}{},
				DNSResolution:       map[string]interface{}{},
				KeyVault:            map[string]interface{}{},
				Nsg:                 map[string]interface{}{},
				Others:              map[string]interface{}{},
				ResourceLock:        map[string]interface{}{},
				StorageAccount:      map[string]interface{}{},
				Udr:                 map[string]interface{}{},
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
	// TODO: use response item
	_ = res
}
