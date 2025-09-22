package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/LogicApps_DeleteDeployWorkflowArtifacts.json
func ExampleLogicAppsClient_DeployWorkflowArtifacts_deleteWorkflowArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewLogicAppsClient().DeployWorkflowArtifacts(ctx, "testrg123", "testapp2", "testapp2", &armappcontainers.LogicAppsClientDeployWorkflowArtifactsOptions{WorkflowArtifacts: &armappcontainers.WorkflowArtifacts{
		FilesToDelete: []*string{
			to.Ptr("test/workflow.json"),
			to.Ptr("test/")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
