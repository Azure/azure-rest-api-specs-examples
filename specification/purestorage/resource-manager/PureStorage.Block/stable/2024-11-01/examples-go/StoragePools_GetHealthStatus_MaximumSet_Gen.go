package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01/StoragePools_GetHealthStatus_MaximumSet_Gen.json
func ExampleStoragePoolsClient_GetHealthStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStoragePoolsClient().GetHealthStatus(ctx, "rgpurestorage", "storagePoolname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.StoragePoolsClientGetHealthStatusResponse{
	// 	StoragePoolHealthInfo: &armpurestorageblock.StoragePoolHealthInfo{
	// 		Health: &armpurestorageblock.HealthDetails{
	// 			UsedCapacityPercentage: to.Ptr[float64](21),
	// 			BandwidthUsage: &armpurestorageblock.BandwidthUsage{
	// 				Current: to.Ptr[int64](15),
	// 				Provisioned: to.Ptr[int64](23),
	// 				Max: to.Ptr[int64](21),
	// 			},
	// 			IopsUsage: &armpurestorageblock.IopsUsage{
	// 				Current: to.Ptr[int64](5),
	// 				Provisioned: to.Ptr[int64](19),
	// 				Max: to.Ptr[int64](15),
	// 			},
	// 			Space: &armpurestorageblock.Space{
	// 				TotalUsed: to.Ptr[int64](28),
	// 				Unique: to.Ptr[int64](4),
	// 				Snapshots: to.Ptr[int64](5),
	// 				Shared: to.Ptr[int64](9),
	// 			},
	// 			DataReductionRatio: to.Ptr[float64](28),
	// 			EstimatedMaxCapacity: to.Ptr[int64](1),
	// 		},
	// 		Alerts: []*armpurestorageblock.Alert{
	// 			{
	// 				Level: to.Ptr(armpurestorageblock.AlertLevelInfo),
	// 				Message: to.Ptr("uasvakccxihgmdz"),
	// 			},
	// 		},
	// 	},
	// }
}
