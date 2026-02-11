package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/ApiKeys_CreateOrUpdate.json
func ExampleAPIKeysClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIKeysClient().CreateOrUpdate(ctx, "myResourceGroup", "myDeployment", "myApiKey", armnginx.DeploymentAPIKeyRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnginx.APIKeysClientCreateOrUpdateResponse{
	// 	DeploymentAPIKeyResponse: &armnginx.DeploymentAPIKeyResponse{
	// 		Name: to.Ptr("myApiKey"),
	// 		Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/apiKeys"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/apiKeys/myApiKey"),
	// 		Properties: &armnginx.DeploymentAPIKeyResponseProperties{
	// 			EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-01T00:00:00Z"); return t}()),
	// 			Hint: to.Ptr("000"),
	// 		},
	// 	},
	// }
}
