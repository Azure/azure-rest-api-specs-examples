package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PutDeploymentAtScope.json
func ExampleDeploymentsClient_BeginCreateOrUpdateAtScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdateAtScope(ctx,
		"providers/Microsoft.Management/managementGroups/my-management-group-id",
		"my-deployment",
		armresources.Deployment{
			Location: to.Ptr("eastus"),
			Properties: &armresources.DeploymentProperties{
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
				Parameters: map[string]interface{}{},
				TemplateLink: &armresources.TemplateLink{
					URI: to.Ptr("https://example.com/exampleTemplate.json"),
				},
			},
			Tags: map[string]*string{
				"tagKey1": to.Ptr("tag-value-1"),
				"tagKey2": to.Ptr("tag-value-2"),
			},
		},
		nil)
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
