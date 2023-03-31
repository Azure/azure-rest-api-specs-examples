package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/DeleteAzurePipeline.json
func ExamplePipelinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPipelinesClient().Delete(ctx, "myAspNetWebAppPipeline-rg", "myAspNetWebAppPipeline", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
