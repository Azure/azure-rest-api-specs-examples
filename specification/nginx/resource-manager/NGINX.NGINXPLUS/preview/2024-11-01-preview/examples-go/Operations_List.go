package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8d26b0e4c1886458fa56c22aac09c3e3e9a5c9e/specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-11-01-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armnginx.OperationListResult{
		// 	Value: []*armnginx.OperationResult{
		// 		{
		// 			Name: to.Ptr("Nginx.NginxPlus/nginxDeployments/write"),
		// 			Display: &armnginx.OperationDisplay{
		// 				Description: to.Ptr("Write deployments resource"),
		// 				Operation: to.Ptr("write"),
		// 				Provider: to.Ptr("Nginx.NginxPlus"),
		// 				Resource: to.Ptr("deployments"),
		// 			},
		// 	}},
		// }
	}
}
