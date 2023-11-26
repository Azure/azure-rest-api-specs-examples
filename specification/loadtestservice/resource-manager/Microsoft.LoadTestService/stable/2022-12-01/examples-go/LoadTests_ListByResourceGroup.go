package armloadtesting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListByResourceGroup.json
func ExampleLoadTestsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armloadtesting.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadTestsClient().NewListByResourceGroupPager("dummyrg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.LoadTestResourcePageList = armloadtesting.LoadTestResourcePageList{
		// 	Value: []*armloadtesting.LoadTestResource{
		// 		{
		// 			Name: to.Ptr("myLoadTest"),
		// 			Type: to.Ptr("Microsoft.LoadTestService/loadTests"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/dummyrg/providers/Microsoft.LoadTestService/loadTests/myLoadTest"),
		// 			SystemData: &armloadtesting.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("userId1001"),
		// 				CreatedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-28T12:32:33.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("userId1001"),
		// 				LastModifiedByType: to.Ptr(armloadtesting.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"Team": to.Ptr("Dev Exp"),
		// 			},
		// 			Properties: &armloadtesting.LoadTestProperties{
		// 				Description: to.Ptr("This is new load test resource"),
		// 				DataPlaneURI: to.Ptr("https://myLoadTest.00000000-0000-0000-0000-000000000000.cnt-dp.domain.com"),
		// 				ProvisioningState: to.Ptr(armloadtesting.ResourceStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
