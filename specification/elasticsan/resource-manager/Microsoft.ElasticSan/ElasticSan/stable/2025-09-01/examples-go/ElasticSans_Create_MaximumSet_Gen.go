package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: 2025-09-01/ElasticSans_Create_MaximumSet_Gen.json
func ExampleElasticSansClient_BeginCreate_elasticSansCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewElasticSansClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", armelasticsan.ElasticSan{
		Location: to.Ptr("France Central"),
		Properties: &armelasticsan.Properties{
			AutoScaleProperties: &armelasticsan.AutoScaleProperties{
				ScaleUpProperties: &armelasticsan.ScaleUpProperties{
					AutoScalePolicyEnforcement:  to.Ptr(armelasticsan.AutoScalePolicyEnforcementNone),
					CapacityUnitScaleUpLimitTiB: to.Ptr[int64](17),
					IncreaseCapacityUnitByTiB:   to.Ptr[int64](4),
					UnusedSizeTiB:               to.Ptr[int64](24),
				},
			},
			AvailabilityZones: []*string{
				to.Ptr("1"),
			},
			BaseSizeTiB:             to.Ptr[int64](5),
			ExtendedCapacitySizeTiB: to.Ptr[int64](25),
			PublicNetworkAccess:     to.Ptr(armelasticsan.PublicNetworkAccessEnabled),
			SKU: &armelasticsan.SKU{
				Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
				Tier: to.Ptr(armelasticsan.SKUTierPremium),
			},
		},
		Tags: map[string]*string{
			"key9316": to.Ptr("ihndtieqibtob"),
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
	// res = armelasticsan.ElasticSansClientCreateResponse{
	// 	ElasticSan: &armelasticsan.ElasticSan{
	// 		Name: to.Ptr("vfoatmakv"),
	// 		Type: to.Ptr("Microsoft.ElasticSan/ElasticSans"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}"),
	// 		Location: to.Ptr("France Central"),
	// 		Properties: &armelasticsan.Properties{
	// 			AutoScaleProperties: &armelasticsan.AutoScaleProperties{
	// 				ScaleUpProperties: &armelasticsan.ScaleUpProperties{
	// 					AutoScalePolicyEnforcement: to.Ptr(armelasticsan.AutoScalePolicyEnforcementNone),
	// 					CapacityUnitScaleUpLimitTiB: to.Ptr[int64](17),
	// 					IncreaseCapacityUnitByTiB: to.Ptr[int64](4),
	// 					UnusedSizeTiB: to.Ptr[int64](24),
	// 				},
	// 			},
	// 			AvailabilityZones: []*string{
	// 				to.Ptr("1"),
	// 			},
	// 			BaseSizeTiB: to.Ptr[int64](15),
	// 			ExtendedCapacitySizeTiB: to.Ptr[int64](6),
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armelasticsan.PublicNetworkAccessEnabled),
	// 			SKU: &armelasticsan.SKU{
	// 				Name: to.Ptr(armelasticsan.SKUNamePremiumLRS),
	// 				Tier: to.Ptr(armelasticsan.SKUTierPremium),
	// 			},
	// 			TotalIops: to.Ptr[int64](22),
	// 			TotalMBps: to.Ptr[int64](4),
	// 			TotalSizeTiB: to.Ptr[int64](27),
	// 			TotalVolumeSizeGiB: to.Ptr[int64](15),
	// 			VolumeGroupCount: to.Ptr[int64](24),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 			CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key5002": to.Ptr("lhag"),
	// 		},
	// 	},
	// }
}
