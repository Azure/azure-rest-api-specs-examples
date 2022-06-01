Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv1.0.1/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmachinelearningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/diagnose.json
func ExampleWorkspacesClient_BeginDiagnose() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewWorkspacesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDiagnose(ctx,
		"workspace-1234",
		"testworkspace",
		&armmachinelearningservices.WorkspacesClientBeginDiagnoseOptions{Parameters: &armmachinelearningservices.DiagnoseWorkspaceParameters{
			Value: &armmachinelearningservices.DiagnoseRequestProperties{
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
```
