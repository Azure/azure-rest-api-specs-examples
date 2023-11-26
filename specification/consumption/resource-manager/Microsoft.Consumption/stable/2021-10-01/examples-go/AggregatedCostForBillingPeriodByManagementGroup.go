package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostForBillingPeriodByManagementGroup.json
func ExampleAggregatedCostClient_GetForBillingPeriodByManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAggregatedCostClient().GetForBillingPeriodByManagementGroup(ctx, "managementGroupForTest", "201807", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroupAggregatedCostResult = armconsumption.ManagementGroupAggregatedCostResult{
	// 	Name: to.Ptr("aggregatedcostId1"),
	// 	Type: to.Ptr("Microsoft.Consumption/aggregatedcost"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/managementGroupForTest/providers/Microsoft.Consumption/aggregatedcostId1"),
	// 	Properties: &armconsumption.ManagementGroupAggregatedCostProperties{
	// 		AzureCharges: to.Ptr[float64](250.9876),
	// 		ChargesBilledSeparately: to.Ptr[float64](120.345),
	// 		Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 			{
	// 				Name: to.Ptr("aggregatedcostId2"),
	// 				Type: to.Ptr("Microsoft.Consumption/aggregatedcost"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/managementGroupChildForTest/providers/Microsoft.Consumption/aggregatedcostId2"),
	// 				Properties: &armconsumption.ManagementGroupAggregatedCostProperties{
	// 					AzureCharges: to.Ptr[float64](150),
	// 					ChargesBilledSeparately: to.Ptr[float64](30.345),
	// 					Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 					},
	// 					Currency: to.Ptr("USD"),
	// 					MarketplaceCharges: to.Ptr[float64](50.786),
	// 					UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-31T00:00:00.000Z"); return t}()),
	// 					UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-01T00:00:00.000Z"); return t}()),
	// 				},
	// 		}},
	// 		Currency: to.Ptr("USD"),
	// 		MarketplaceCharges: to.Ptr[float64](150.786),
	// 		UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-31T00:00:00.000Z"); return t}()),
	// 		UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-07-01T00:00:00.000Z"); return t}()),
	// 	},
	// }
}
