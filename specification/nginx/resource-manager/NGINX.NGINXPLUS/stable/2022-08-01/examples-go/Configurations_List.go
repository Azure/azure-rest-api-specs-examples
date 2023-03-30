package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/nginx/resource-manager/NGINX.NGINXPLUS/stable/2022-08-01/examples/Configurations_List.json
func ExampleConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ConfigurationListResponse = armnginx.ConfigurationListResponse{
		// 	Value: []*armnginx.Configuration{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("nginx.nginxplus/nginxDeployments/configurations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Nginx.NginxPlus/nginxDeployments/myDeployment/configurations/default"),
		// 			Properties: &armnginx.ConfigurationProperties{
		// 				Files: []*armnginx.ConfigurationFile{
		// 					{
		// 						Content: to.Ptr("ABCDEF=="),
		// 						VirtualPath: to.Ptr("/etc/nginx/nginx.conf"),
		// 				}},
		// 				Package: &armnginx.ConfigurationPackage{
		// 				},
		// 				ProvisioningState: to.Ptr(armnginx.ProvisioningStateSucceeded),
		// 				RootFile: to.Ptr("/etc/nginx/nginx.conf"),
		// 			},
		// 	}},
		// }
	}
}
