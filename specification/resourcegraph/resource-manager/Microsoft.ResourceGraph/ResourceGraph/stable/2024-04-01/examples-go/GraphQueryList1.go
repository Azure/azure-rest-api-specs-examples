package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2024-04-01/GraphQueryList1.json
func ExampleGraphQueryClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGraphQueryClient("024e2271-06fa-46b6-9079-f1ed3c7b070e").NewListPager("my-resource-group", nil)
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
		// page = armresourcegraph.GraphQueryClientListResponse{
		// 	GraphQueryListResult: armresourcegraph.GraphQueryListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subscriptionId/providers/Microsoft.ResourceGraph/queries?api-version=2019-09-01-preview&$top=10&$skiptoken=skiptoken"),
		// 		Value: []*armresourcegraph.GraphQueryResource{
		// 			{
		// 				Name: to.Ptr("MyDockerVMs"),
		// 				Type: to.Ptr("Microsoft.ResourceGraph/queries"),
		// 				Etag: to.Ptr("5d64408e-4ca3-41f7-b725-6914f3012afa"),
		// 				ID: to.Ptr("/subscriptions/87f4f8b0-83c1-4aa9-b318-5237aeb15264/resourceGroups/my-resource-group/providers/Microsoft.ResourceGraph/queries/MyDockerVMs"),
		// 				Properties: &armresourcegraph.GraphQueryProperties{
		// 					Description: to.Ptr("Docker VMs in PROD"),
		// 					Query: to.Ptr("where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'docker'"),
		// 					ResultKind: to.Ptr(armresourcegraph.ResultKindBasic),
		// 					TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-12T13:51:13-07:00"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("MyTestVMs"),
		// 				Type: to.Ptr("Microsoft.ResourceGraph/queries"),
		// 				Etag: to.Ptr("b0809832-ca62-4133-8f13-0c46580f9db1"),
		// 				ID: to.Ptr("/subscriptions/6abb2f31-3e6a-4134-9968-219a596012a0/resourceGroups/my-resource-group/providers/Microsoft.ResourceGraph/queries/MyTestVMs"),
		// 				Properties: &armresourcegraph.GraphQueryProperties{
		// 					Description: to.Ptr("Test VMs in PROD"),
		// 					Query: to.Ptr("where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'test'"),
		// 					ResultKind: to.Ptr(armresourcegraph.ResultKindBasic),
		// 					TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-13T13:51:13-07:00"); return t}()),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
