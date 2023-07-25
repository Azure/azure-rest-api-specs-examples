package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/WebPubSubReplicas_List.json
func ExampleReplicasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListPager("myResourceGroup", "myWebPubSubService", nil)
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
		// page.ReplicaList = armwebpubsub.ReplicaList{
		// 	Value: []*armwebpubsub.Replica{
		// 		{
		// 			Name: to.Ptr("myWebPubSubService-eastus"),
		// 			Type: to.Ptr("Microsoft.SignalRService/WebPubSub"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/replicas/myWebPubSubService-eastus"),
		// 			SystemData: &armwebpubsub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armwebpubsub.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armwebpubsub.ReplicaProperties{
		// 				ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armwebpubsub.ResourceSKU{
		// 				Name: to.Ptr("Premium_P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierPremium),
		// 			},
		// 	}},
		// }
	}
}
