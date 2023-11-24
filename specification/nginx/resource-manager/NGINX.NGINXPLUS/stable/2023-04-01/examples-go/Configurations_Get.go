package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2023-04-01/examples/Configurations_Get.json
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
	// res.Configuration = armnginx.Configuration{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("nginx.nginxplus/nginxDeployments/configurations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/configurations/default"),
	// 	Properties: &armnginx.ConfigurationProperties{
	// 		Files: []*armnginx.ConfigurationFile{
	// 			{
	// 				Content: to.Ptr("ABCDEF=="),
	// 				VirtualPath: to.Ptr("/etc/nginx/nginx.conf"),
	// 		}},
	// 		Package: &armnginx.ConfigurationPackage{
	// 		},
	// 		ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
	// 		RootFile: to.Ptr("/etc/nginx/nginx.conf"),
	// 	},
	// }
}
