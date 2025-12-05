package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticCapacityPools_ChangeZone.json
func ExampleElasticCapacityPoolsClient_BeginChangeZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticCapacityPoolsClient().BeginChangeZone(ctx, "myRG", "account1", "pool1", armnetapp.ChangeZoneRequest{
		NewZone: to.Ptr("3"),
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
	// res = armnetapp.ElasticCapacityPoolsClientChangeZoneResponse{
	// 	ElasticCapacityPool: &armnetapp.ElasticCapacityPool{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1"),
	// 		Name: to.Ptr("account1/pool1"),
	// 		Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticCapacityPools"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.ElasticCapacityPoolProperties{
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			ServiceLevel: to.Ptr(armnetapp.ElasticServiceLevelZoneRedundant),
	// 			Size: to.Ptr[int64](4398046511104),
	// 			TotalThroughputMibps: to.Ptr[float64](281.474),
	// 			CurrentZone: to.Ptr("3"),
	// 			AvailabilityStatus: to.Ptr(armnetapp.ElasticResourceAvailabilityStatusOnline),
	// 			SubnetResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
	// 		},
	// 	},
	// }
}
