package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a7af6049f4b4743ef3b649f3852bcc7bd9a43ee0/specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/FarmBeatsModels_ListByResourceGroup.json
func ExampleFarmBeatsModelsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagrifood.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFarmBeatsModelsClient().NewListByResourceGroupPager("examples-rg", &armagrifood.FarmBeatsModelsClientListByResourceGroupOptions{MaxPageSize: nil,
		SkipToken: nil,
	})
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
		// page.FarmBeatsListResponse = armagrifood.FarmBeatsListResponse{
		// 	Value: []*armagrifood.FarmBeats{
		// 		{
		// 			Name: to.Ptr("examples-farmBeatsResourceName"),
		// 			Type: to.Ptr("Microsoft.AgFoodPlatform/farmBeats"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/Microsoft.AgFoodPlatform/farmBeats/examples-farmbeatsResourceName"),
		// 			SystemData: &armagrifood.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armagrifood.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armagrifood.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus2"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armagrifood.FarmBeatsProperties{
		// 				InstanceURI: to.Ptr("https://examples-farmbeatsResourceName.eastus2.farmbeats.azure.net"),
		// 				ProvisioningState: to.Ptr(armagrifood.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
