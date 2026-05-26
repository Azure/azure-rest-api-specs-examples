package armresourcegraph_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// Generated from example definition: 2024-04-01/GraphQueryGet.json
func ExampleGraphQueryClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcegraph.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGraphQueryClient("024e2271-06fa-46b6-9079-f1ed3c7b070e").Get(ctx, "my-resource-group", "MyDockerVMs", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourcegraph.GraphQueryClientGetResponse{
	// 	GraphQueryResource: armresourcegraph.GraphQueryResource{
	// 		Name: to.Ptr("MyDockerVMs"),
	// 		Type: to.Ptr("Microsoft.ResourceGraph/queries"),
	// 		Etag: to.Ptr("5d64408e-4ca3-41f7-b725-6914f3012afa"),
	// 		ID: to.Ptr("/subscriptions/024e2271-06fa-46b6-9079-f1ed3c7b070e/resourceGroups/my-resource-group/providers/Microsoft.ResourceGraph/queries/MyDockerVMs"),
	// 		Properties: &armresourcegraph.GraphQueryProperties{
	// 			Description: to.Ptr("Docker VMs in PROD"),
	// 			Query: to.Ptr("where isnotnull(tags['Prod']) and properties.extensions[0].Name == 'docker'"),
	// 			ResultKind: to.Ptr(armresourcegraph.ResultKindBasic),
	// 			TimeModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-12T13:51:13-07:00"); return t}()),
	// 		},
	// 	},
	// }
}
