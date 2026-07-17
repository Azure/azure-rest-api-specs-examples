package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption/v2"
)

// Generated from example definition: 2024-08-01/AggregatedCostByManagementGroup.json
func ExampleAggregatedCostClient_GetByManagementGroup_aggregatedCostByManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAggregatedCostClient().GetByManagementGroup(ctx, "managementGroupForTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconsumption.AggregatedCostClientGetByManagementGroupResponse{
	// 	ManagementGroupAggregatedCostResult: armconsumption.ManagementGroupAggregatedCostResult{
	// 		Name: to.Ptr("aggregatedcostId1"),
	// 		Type: to.Ptr("Microsoft.Consumption/aggregatedcost"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/managementGroupForTest/providers/Microsoft.Consumption/aggregatedcostId1"),
	// 		Properties: &armconsumption.ManagementGroupAggregatedCostProperties{
	// 			AzureCharges: to.Ptr[float64](250.9876),
	// 			ChargesBilledSeparately: to.Ptr[float64](120.345),
	// 			Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 				{
	// 					Name: to.Ptr("aggregatedcostId2"),
	// 					Type: to.Ptr("Microsoft.Consumption/aggregatedcost"),
	// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/managementGroupChildForTest/providers/Microsoft.Consumption/aggregatedcostId2"),
	// 					Properties: &armconsumption.ManagementGroupAggregatedCostProperties{
	// 						AzureCharges: to.Ptr[float64](150),
	// 						ChargesBilledSeparately: to.Ptr[float64](30.345),
	// 						Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 						},
	// 						Currency: to.Ptr("USD"),
	// 						ExcludedSubscriptions: []*string{
	// 						},
	// 						IncludedSubscriptions: []*string{
	// 							to.Ptr("c349567d-c83a-48c9-ab0e-578c69dc97a4"),
	// 						},
	// 						MarketplaceCharges: to.Ptr[float64](50.786),
	// 						UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-01T00:00:00.0000000Z"); return t}()),
	// 						UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T00:00:00.0000000Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			Currency: to.Ptr("USD"),
	// 			ExcludedSubscriptions: []*string{
	// 			},
	// 			IncludedSubscriptions: []*string{
	// 				to.Ptr("1caaa5a3-2b66-438e-8ab4-bce37d518c5d"),
	// 			},
	// 			MarketplaceCharges: to.Ptr[float64](150.786),
	// 			UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-01T00:00:00.0000000Z"); return t}()),
	// 			UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-03-01T00:00:00.0000000Z"); return t}()),
	// 		},
	// 	},
	// }
}
