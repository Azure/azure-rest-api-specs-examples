package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBGraphAPIComputeServiceGet.json
func ExampleServiceClient_Get_graphApiComputeServiceGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceClient().Get(ctx, "rg1", "ddb1", "GraphAPICompute", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.ServiceClientGetResponse{
	// 	ServiceResource: armcosmos.ServiceResource{
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/GraphAPICompute"),
	// 		Name: to.Ptr("GraphAPICompute"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 		Properties: &armcosmos.GraphAPIComputeServiceResourceProperties{
	// 			Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 			InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 			InstanceCount: to.Ptr[int32](1),
	// 			ServiceType: to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
	// 			GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute.gremlin.cosmos.windows-int.net/"),
	// 			Locations: []*armcosmos.GraphAPIComputeRegionalServiceResource{
	// 				{
	// 					Name: to.Ptr("GraphAPICompute-westus2"),
	// 					Location: to.Ptr("West US 2"),
	// 					Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 					GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute-westus.gremlin.cosmos.windows-int.net/"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
