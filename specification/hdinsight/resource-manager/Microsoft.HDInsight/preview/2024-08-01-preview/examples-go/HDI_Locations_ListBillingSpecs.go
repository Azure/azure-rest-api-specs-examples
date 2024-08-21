package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Locations_ListBillingSpecs.json
func ExampleLocationsClient_ListBillingSpecs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().ListBillingSpecs(ctx, "East US 2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BillingResponseListResult = armhdinsight.BillingResponseListResult{
	// 	BillingResources: []*armhdinsight.BillingResources{
	// 		{
	// 			BillingMeters: []*armhdinsight.BillingMeters{
	// 			},
	// 			DiskBillingMeters: []*armhdinsight.DiskBillingMeters{
	// 			},
	// 			Region: to.Ptr("East US 2"),
	// 		},
	// 		{
	// 			BillingMeters: []*armhdinsight.BillingMeters{
	// 				{
	// 					Meter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					MeterParameter: to.Ptr("default"),
	// 					Unit: to.Ptr("CoreHours"),
	// 				},
	// 				{
	// 					Meter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					MeterParameter: to.Ptr("Kafka"),
	// 					Unit: to.Ptr("CoreHours"),
	// 			}},
	// 			DiskBillingMeters: []*armhdinsight.DiskBillingMeters{
	// 				{
	// 					DiskRpMeter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					SKU: to.Ptr("All"),
	// 					Tier: to.Ptr(armhdinsight.TierStandard),
	// 				},
	// 				{
	// 					DiskRpMeter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					SKU: to.Ptr("All"),
	// 					Tier: to.Ptr(armhdinsight.TierStandard),
	// 			}},
	// 			Region: to.Ptr("default"),
	// 	}},
	// 	VMSizeFilters: []*armhdinsight.VMSizeCompatibilityFilterV2{
	// 		{
	// 			FilterMode: to.Ptr(armhdinsight.FilterModeExclude),
	// 	}},
	// 	VMSizes: []*string{
	// 		to.Ptr("A5"),
	// 		to.Ptr("A6"),
	// 		to.Ptr("A7")},
	// 	}
}
