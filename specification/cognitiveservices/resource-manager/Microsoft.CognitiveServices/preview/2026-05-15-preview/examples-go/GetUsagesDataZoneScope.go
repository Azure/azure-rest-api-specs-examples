package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/GetUsagesDataZoneScope.json
func ExampleAccountsClient_ListUsages_getUsagesDataZoneScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("5a4f5c2e-6983-4ccb-bd34-2196d5b5bbd3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListUsages(ctx, "myResourceGroup", "TestUsage02", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AccountsClientListUsagesResponse{
	// 	UsageListResult: armcognitiveservices.UsageListResult{
	// 		Value: []*armcognitiveservices.Usage{
	// 			{
	// 				Name: &armcognitiveservices.MetricName{
	// 					LocalizedValue: to.Ptr("Face.Transactions"),
	// 					Value: to.Ptr("Face.Transactions"),
	// 				},
	// 				CurrentValue: to.Ptr[float64](4),
	// 				Limit: to.Ptr[float64](40000),
	// 				NextResetTime: to.Ptr("2018-03-28T09:33:51Z"),
	// 				QuotaPeriod: to.Ptr("30.00:00:00"),
	// 				Status: to.Ptr(armcognitiveservices.QuotaUsageStatusIncluded),
	// 				Unit: to.Ptr(armcognitiveservices.UnitTypeCount),
	// 				ScopeType: to.Ptr(armcognitiveservices.QuotaScopeTypeDataZone),
	// 				ScopeID: to.Ptr("US"),
	// 			},
	// 		},
	// 	},
	// }
}
