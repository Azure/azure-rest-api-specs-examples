package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/402006d2796cdd3894d013d83e77b46a5c844005/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-11-15/examples/CosmosDBGraphAPIComputeServiceCreate.json
func ExampleServiceClient_BeginCreate_graphApiComputeServiceCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewServiceClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "rg1", "ddb1", "GraphAPICompute", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armcosmos.ServiceResource{
	// 	Name: to.Ptr("GraphAPICompute"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/services"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/services/GraphAPICompute"),
	// 	Properties: &armcosmos.GraphAPIComputeServiceResourceProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-25T12:56:05.4622517Z"); return t}()),
	// 		InstanceCount: to.Ptr[int32](1),
	// 		InstanceSize: to.Ptr(armcosmos.ServiceSizeCosmosD4S),
	// 		ServiceType: to.Ptr(armcosmos.ServiceTypeGraphAPICompute),
	// 		Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 		GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute.gremlin.cosmos.windows-int.net/"),
	// 		Locations: []*armcosmos.GraphAPIComputeRegionalServiceResource{
	// 			{
	// 				Name: to.Ptr("GraphAPICompute-westus2"),
	// 				Location: to.Ptr("West US 2"),
	// 				Status: to.Ptr(armcosmos.ServiceStatusRunning),
	// 				GraphAPIComputeEndpoint: to.Ptr("https://graphAPICompute-westus.gremlin.cosmos.windows-int.net/"),
	// 		}},
	// 	},
	// }
}
