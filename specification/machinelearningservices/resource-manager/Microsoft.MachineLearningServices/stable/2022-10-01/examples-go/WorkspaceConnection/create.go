package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/WorkspaceConnection/create.json
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
	res, err := client.Create(ctx, "resourceGroup-1", "workspace-1", "connection-1", armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
		Properties: &armmachinelearning.NoneAuthTypeWorkspaceConnectionProperties{
			AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeNone),
			Category: to.Ptr(armmachinelearning.ConnectionCategoryContainerRegistry),
			Target:   to.Ptr("www.facebook.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
