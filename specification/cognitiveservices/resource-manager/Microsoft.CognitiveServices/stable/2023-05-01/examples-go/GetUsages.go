package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetUsages.json
func ExampleAccountsClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListUsages(ctx, "myResourceGroup", "TestUsage02", &armcognitiveservices.AccountsClientListUsagesOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UsageListResult = armcognitiveservices.UsageListResult{
	// 	Value: []*armcognitiveservices.Usage{
	// 		{
	// 			Name: &armcognitiveservices.MetricName{
	// 				LocalizedValue: to.Ptr("Face.Transactions"),
	// 				Value: to.Ptr("Face.Transactions"),
	// 			},
	// 			CurrentValue: to.Ptr[float64](3),
	// 			Limit: to.Ptr[float64](30000),
	// 			NextResetTime: to.Ptr("2018-03-28T09:33:51Z"),
	// 			QuotaPeriod: to.Ptr("30.00:00:00"),
	// 			Status: to.Ptr(armcognitiveservices.QuotaUsageStatusIncluded),
	// 			Unit: to.Ptr(armcognitiveservices.UnitTypeCount),
	// 	}},
	// }
}
