package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/StoragePools_GetHealthStatus_MaximumSet_Gen.json
func ExampleStoragePoolsClient_GetHealthStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStoragePoolsClient().GetHealthStatus(ctx, "rgpurestorage", "storagepool-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpurestorageblock.StoragePoolsClientGetHealthStatusResponse{
	// 	StoragePoolHealthInfo: armpurestorageblock.StoragePoolHealthInfo{
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
	// 				TotalUsed: to.Ptr[int64](1073741824),
	// 				Unique: to.Ptr[int64](268435456),
	// 				Snapshots: to.Ptr[int64](53687091236870910),
	// 				Shared: to.Ptr[int64](268435456),
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
