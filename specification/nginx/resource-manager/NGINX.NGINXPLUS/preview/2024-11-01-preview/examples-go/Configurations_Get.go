package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8d26b0e4c1886458fa56c22aac09c3e3e9a5c9e/specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Configurations_Get.json
func ExampleConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().Get(ctx, "myResourceGroup", "myDeployment", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationResponse = armnginx.ConfigurationResponse{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("nginx.nginxplus/nginxDeployments/configurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/configurations/default"),
	// 	Properties: &armnginx.ConfigurationResponseProperties{
	// 		Files: []*armnginx.ConfigurationFile{
	// 			{
	// 				Content: to.Ptr("ABCDEF=="),
	// 				VirtualPath: to.Ptr("/etc/nginx/nginx.conf"),
	// 		}},
	// 		Package: &armnginx.ConfigurationPackage{
	// 		},
	// 		ProtectedFiles: []*armnginx.ConfigurationProtectedFileResponse{
	// 			{
	// 				ContentHash: to.Ptr("sha256:1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"),
	// 				VirtualPath: to.Ptr("/etc/nginx/protected-file.cert"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 		RootFile: to.Ptr("/etc/nginx/nginx.conf"),
	// 	},
	// }
}
