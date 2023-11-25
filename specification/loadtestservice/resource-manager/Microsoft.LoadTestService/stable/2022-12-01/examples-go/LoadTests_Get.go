package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Get.json
func ExampleLoadTestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadTestsClient().Get(ctx, "dummyrg", "myLoadTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoadTestResource = armloadtesting.LoadTestResource{
	// 	Name: to.Ptr("myLoadTest"),
	// 	Type: to.Ptr("Microsoft.LoadTestService/loadTests"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.LoadTestService/loadTests/myLoadTest"),
	// 	SystemData: &armloadtesting.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("userId1001"),
	// 		CreatedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("userId1001"),
	// 		LastModifiedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"Team": to.Ptr("Dev Exp"),
	// 	},
	// 	Properties: &armloadtesting.LoadTestProperties{
	// 		Description: to.Ptr("This is new load test resource"),
	// 		DataPlaneURI: to.Ptr("https://myLoadTest.00000000-0000-0000-0000-000000000000.cnt-dp.domain.com"),
	// 		ProvisioningState: to.Ptr(armloadtesting.ResourceStateSucceeded),
	// 	},
	// }
}
