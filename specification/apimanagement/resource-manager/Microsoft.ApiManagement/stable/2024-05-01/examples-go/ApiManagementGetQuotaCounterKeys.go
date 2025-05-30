package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetQuotaCounterKeys.json
func ExampleQuotaByCounterKeysClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotaByCounterKeysClient().ListByService(ctx, "rg1", "apimService1", "ba", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QuotaCounterCollection = armapimanagement.QuotaCounterCollection{
	// 	Value: []*armapimanagement.QuotaCounterContract{
	// 		{
	// 			CounterKey: to.Ptr("ba"),
	// 			PeriodEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-08T16:54:40.000Z"); return t}()),
	// 			PeriodKey: to.Ptr("0_P3Y6M4DT12H30M5S"),
	// 			PeriodStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-08-04T04:24:35.000Z"); return t}()),
	// 			Value: &armapimanagement.QuotaCounterValueContractProperties{
	// 				CallsCount: to.Ptr[int32](5),
	// 				KbTransferred: to.Ptr[float64](2.5830078125),
	// 			},
	// 	}},
	// }
}
