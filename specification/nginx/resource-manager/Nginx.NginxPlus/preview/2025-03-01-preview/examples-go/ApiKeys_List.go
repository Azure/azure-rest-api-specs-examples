package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/ApiKeys_List.json
func ExampleAPIKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPIKeysClient().NewListPager("myResourceGroup", "myDeployment", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnginx.APIKeysClientListResponse{
		// 	DeploymentAPIKeyListResponse: armnginx.DeploymentAPIKeyListResponse{
		// 		Value: []*armnginx.DeploymentAPIKeyResponse{
		// 			{
		// 				Name: to.Ptr("myApiKey"),
		// 				Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/apiKeys"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/apiKeys/myApiKey"),
		// 				Properties: &armnginx.DeploymentAPIKeyResponseProperties{
		// 					EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-01T00:00:00Z"); return t}()),
		// 					Hint: to.Ptr("000"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myApiKey2"),
		// 				Type: to.Ptr("Nginx.NginxPlus/nginxDeployments/apiKeys"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/apiKeys/myApiKey2"),
		// 				Properties: &armnginx.DeploymentAPIKeyResponseProperties{
		// 					EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-01T00:00:00Z"); return t}()),
		// 					Hint: to.Ptr("111"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
