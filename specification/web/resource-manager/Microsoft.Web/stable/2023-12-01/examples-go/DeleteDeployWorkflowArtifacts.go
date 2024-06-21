package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/de1f3772629b6f4d3ac01548a5f6d719bfb97c9e/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/DeleteDeployWorkflowArtifacts.json
func ExampleWebAppsClient_DeployWorkflowArtifacts_deleteWorkflowArtifacts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWebAppsClient().DeployWorkflowArtifacts(ctx, "testrg123", "testsite2", &armappservice.WebAppsClientDeployWorkflowArtifactsOptions{WorkflowArtifacts: &armappservice.WorkflowArtifacts{
		FilesToDelete: []*string{
			to.Ptr("test/workflow.json"),
			to.Ptr("test/")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
