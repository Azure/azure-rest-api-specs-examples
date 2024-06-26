package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ConsumerSourceDataSets_ListByShareSubscription.json
func ExampleConsumerSourceDataSetsClient_NewListByShareSubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatashare.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConsumerSourceDataSetsClient().NewListByShareSubscriptionPager("SampleResourceGroup", "Account1", "Share1", &armdatashare.ConsumerSourceDataSetsClientListByShareSubscriptionOptions{SkipToken: nil})
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
		// page.ConsumerSourceDataSetList = armdatashare.ConsumerSourceDataSetList{
		// 	Value: []*armdatashare.ConsumerSourceDataSet{
		// 		{
		// 			Name: to.Ptr("invitation1"),
		// 			Type: to.Ptr("Microsoft.DataShare/accounts/sharesubscriptions/consumerSourceDataSets"),
		// 			ID: to.Ptr("/subscriptions/433a8dfd-e5d5-4e77-ad86-90acdc75eb1a/resourceGroups/SampleResourceGroup/providers/Microsoft.DataShare/accounts/Account1/shareSubscriptions/ShareSubscription1/consumerSourceDataSets/4256e2cf-0f82-4865-961b-12f83333f487"),
		// 			Properties: &armdatashare.ConsumerSourceDataSetProperties{
		// 				DataSetID: to.Ptr("0b9d4394-8bb3-49a1-aa4f-4be49cd10375"),
		// 				DataSetName: to.Ptr("input.text"),
		// 				DataSetPath: to.Ptr("cars.text"),
		// 				DataSetType: to.Ptr(armdatashare.DataSetTypeBlob),
		// 			},
		// 	}},
		// }
	}
}
