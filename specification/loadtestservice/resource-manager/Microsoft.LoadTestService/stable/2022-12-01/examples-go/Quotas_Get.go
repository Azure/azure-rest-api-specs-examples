package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_Get.json
func ExampleQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQuotasClient().Get(ctx, "westus", "testQuotaBucket", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QuotaResource = armloadtesting.QuotaResource{
	// 	Name: to.Ptr("testQuotaBucket"),
	// 	Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
	// 	Properties: &armloadtesting.QuotaResourceProperties{
	// 		Limit: to.Ptr[int32](50),
	// 		Usage: to.Ptr[int32](20),
	// 	},
	// }
}
