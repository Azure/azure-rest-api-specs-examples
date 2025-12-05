package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Pools_CreateOrUpdate.json
func ExamplePoolsClient_BeginCreateOrUpdate_poolsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPoolsClient().BeginCreateOrUpdate(ctx, "myRG", "account1", "pool1", armnetapp.CapacityPool{
		Location: to.Ptr("eastus"),
		Properties: &armnetapp.PoolProperties{
			QosType:      to.Ptr(armnetapp.QosTypeAuto),
			ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
			Size:         to.Ptr[int64](4398046511104),
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
	// res = armnetapp.PoolsClientCreateOrUpdateResponse{
	// 	CapacityPool: &armnetapp.CapacityPool{
	// 		Name: to.Ptr("account1/pool1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetapp.PoolProperties{
	// 			PoolID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			QosType: to.Ptr(armnetapp.QosTypeAuto),
	// 			ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
	// 			Size: to.Ptr[int64](4398046511104),
	// 			TotalThroughputMibps: to.Ptr[float32](281.474),
	// 		},
	// 	},
	// }
}
