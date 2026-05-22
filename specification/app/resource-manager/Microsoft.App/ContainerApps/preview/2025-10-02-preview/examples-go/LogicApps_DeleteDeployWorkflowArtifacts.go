package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/LogicApps_DeleteDeployWorkflowArtifacts.json
func ExampleLogicAppsClient_DeployWorkflowArtifacts_deleteWorkflowArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogicAppsClient().DeployWorkflowArtifacts(ctx, "testrg123", "testapp2", "testapp2", &armappcontainers.LogicAppsClientDeployWorkflowArtifactsOptions{
		WorkflowArtifacts: &armappcontainers.WorkflowArtifacts{
			FilesToDelete: []*string{
				to.Ptr("test/workflow.json"),
				to.Ptr("test/"),
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.LogicAppsClientDeployWorkflowArtifactsResponse{
	// }
}
