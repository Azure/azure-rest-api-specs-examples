package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/GitHubOAuthCallback.json
func ExampleDeveloperHubServiceClient_GitHubOAuthCallback() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevhub.NewDeveloperHubServiceClient("subscriptionId1", "3584d83530557fdd1f46af8289938c8ef79f9dc5", "12345678-3456-7890-5678-012345678901", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GitHubOAuthCallback(ctx, "eastus2euap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
