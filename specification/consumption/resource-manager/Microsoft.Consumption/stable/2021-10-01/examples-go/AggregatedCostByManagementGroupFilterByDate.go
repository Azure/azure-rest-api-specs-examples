package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroupFilterByDate.json
func ExampleAggregatedCostClient_GetByManagementGroup_aggregatedCostByManagementGroupFilterByDate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconsumption.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAggregatedCostClient().GetByManagementGroup(ctx, "managementGroupForTest", &armconsumption.AggregatedCostClientGetByManagementGroupOptions{Filter: to.Ptr("usageStart ge '2018-08-15' and properties/usageStart le '2018-08-31'")})
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
	// 		AzureCharges: to.Ptr[float64](150.9876),
	// 		ChargesBilledSeparately: to.Ptr[float64](90.345),
	// 		Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 			{
	// 				Name: to.Ptr("aggregatedcostId2"),
	// 				Type: to.Ptr("Microsoft.Consumption/aggregatedcost"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/managementGroupChildForTest/providers/Microsoft.Consumption/aggregatedcostId2"),
	// 				Properties: &armconsumption.ManagementGroupAggregatedCostProperties{
	// 					AzureCharges: to.Ptr[float64](50),
	// 					ChargesBilledSeparately: to.Ptr[float64](30.345),
	// 					Children: []*armconsumption.ManagementGroupAggregatedCostResult{
	// 					},
	// 					Currency: to.Ptr("USD"),
	// 					ExcludedSubscriptions: []*string{
	// 					},
	// 					IncludedSubscriptions: []*string{
	// 						to.Ptr("c349567d-c83a-48c9-ab0e-578c69dc97a4")},
	// 						MarketplaceCharges: to.Ptr[float64](10.786),
	// 						UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-31T00:00:00.000Z"); return t}()),
	// 						UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T00:00:00.000Z"); return t}()),
	// 					},
	// 			}},
	// 			Currency: to.Ptr("USD"),
	// 			ExcludedSubscriptions: []*string{
	// 			},
	// 			IncludedSubscriptions: []*string{
	// 				to.Ptr("1caaa5a3-2b66-438e-8ab4-bce37d518c5d")},
	// 				MarketplaceCharges: to.Ptr[float64](80.786),
	// 				UsageEnd: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-31T00:00:00.000Z"); return t}()),
	// 				UsageStart: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T00:00:00.000Z"); return t}()),
	// 			},
	// 		}
}
