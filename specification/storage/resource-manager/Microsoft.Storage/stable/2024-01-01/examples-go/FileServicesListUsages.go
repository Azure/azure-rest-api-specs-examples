package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileServicesListUsages.json
func ExampleFileServicesClient_NewListServiceUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileServicesClient().NewListServiceUsagesPager("res4410", "sto8607", &armstorage.FileServicesClientListServiceUsagesOptions{Maxpagesize: nil})
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
		// page.FileServiceUsages = armstorage.FileServiceUsages{
		// 	Value: []*armstorage.FileServiceUsage{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/fileServices/usages"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/res4410/providers/Microsoft.Storage/storageAccounts/sto8607/fileServices/default/usages/default"),
		// 			Properties: &armstorage.FileServiceUsageProperties{
		// 				BurstingConstants: &armstorage.BurstingConstants{
		// 					BurstFloorIOPS: to.Ptr[int32](10000),
		// 					BurstIOScalar: to.Ptr[float64](3),
		// 					BurstTimeframeSeconds: to.Ptr[int32](3600),
		// 				},
		// 				FileShareLimits: &armstorage.FileShareLimits{
		// 					MaxProvisionedBandwidthMiBPerSec: to.Ptr[int32](10340),
		// 					MaxProvisionedIOPS: to.Ptr[int32](102400),
		// 					MaxProvisionedStorageGiB: to.Ptr[int32](262144),
		// 					MinProvisionedBandwidthMiBPerSec: to.Ptr[int32](125),
		// 					MinProvisionedIOPS: to.Ptr[int32](3000),
		// 					MinProvisionedStorageGiB: to.Ptr[int32](32),
		// 				},
		// 				FileShareRecommendations: &armstorage.FileShareRecommendations{
		// 					BandwidthScalar: to.Ptr[float64](0.1),
		// 					BaseBandwidthMiBPerSec: to.Ptr[int32](125),
		// 					BaseIOPS: to.Ptr[int32](3000),
		// 					IoScalar: to.Ptr[float64](1),
		// 				},
		// 				StorageAccountLimits: &armstorage.AccountLimits{
		// 					MaxFileShares: to.Ptr[int32](50),
		// 					MaxProvisionedBandwidthMiBPerSec: to.Ptr[int32](10340),
		// 					MaxProvisionedIOPS: to.Ptr[int32](102400),
		// 					MaxProvisionedStorageGiB: to.Ptr[int32](262144),
		// 				},
		// 				StorageAccountUsage: &armstorage.AccountUsage{
		// 					LiveShares: &armstorage.AccountUsageElements{
		// 						FileShareCount: to.Ptr[int32](2),
		// 						ProvisionedBandwidthMiBPerSec: to.Ptr[int32](258),
		// 						ProvisionedIOPS: to.Ptr[int32](6064),
		// 						ProvisionedStorageGiB: to.Ptr[int32](64),
		// 					},
		// 					SoftDeletedShares: &armstorage.AccountUsageElements{
		// 						FileShareCount: to.Ptr[int32](1),
		// 						ProvisionedBandwidthMiBPerSec: to.Ptr[int32](125),
		// 						ProvisionedIOPS: to.Ptr[int32](3000),
		// 						ProvisionedStorageGiB: to.Ptr[int32](100),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
