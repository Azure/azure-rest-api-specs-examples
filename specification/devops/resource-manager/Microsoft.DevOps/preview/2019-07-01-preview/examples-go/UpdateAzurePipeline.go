package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/UpdateAzurePipeline.json
func ExamplePipelinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		armdevops.PipelineUpdateParameters{
			Tags: map[string]*string{
				"tagKey": to.Ptr("tagvalue"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
