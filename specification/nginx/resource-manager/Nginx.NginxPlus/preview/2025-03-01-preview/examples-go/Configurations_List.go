package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/Configurations_List.json
func ExampleConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConfigurationsClient().NewListPager("myResourceGroup", "myDeployment", nil)
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
		// page = armnginx.ConfigurationsClientListResponse{
		// 	ConfigurationListResponse: armnginx.ConfigurationListResponse{
		// 		Value: []*armnginx.ConfigurationResponse{
		// 			{
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("nginx.nginxplus/nginxDeployments/configurations"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/configurations/default"),
		// 				Properties: &armnginx.ConfigurationResponseProperties{
		// 					Files: []*armnginx.ConfigurationFile{
		// 						{
		// 							Content: to.Ptr("ABCDEF=="),
		// 							VirtualPath: to.Ptr("/etc/nginx/nginx.conf"),
		// 						},
		// 					},
		// 					Package: &armnginx.ConfigurationPackage{
		// 					},
		// 					ProtectedFiles: []*armnginx.ConfigurationProtectedFileResponse{
		// 						{
		// 							ContentHash: to.Ptr("sha256:1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF"),
		// 							VirtualPath: to.Ptr("/etc/nginx/protected-file.cert"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 					RootFile: to.Ptr("/etc/nginx/nginx.conf"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
