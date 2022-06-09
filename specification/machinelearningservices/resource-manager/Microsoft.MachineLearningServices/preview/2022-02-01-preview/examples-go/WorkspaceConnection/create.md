```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/WorkspaceConnection/create.json
func ExampleWorkspaceConnectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewWorkspaceConnectionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"resourceGroup-1",
		"workspace-1",
		"connection-1",
		armmachinelearning.WorkspaceConnection{
			Properties: &armmachinelearning.WorkspaceConnectionProps{
				AuthType: to.Ptr("PAT"),
				Category: to.Ptr("ACR"),
				Target:   to.Ptr("www.facebook.com"),
				Value:    to.Ptr("secrets"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0-beta.1/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
