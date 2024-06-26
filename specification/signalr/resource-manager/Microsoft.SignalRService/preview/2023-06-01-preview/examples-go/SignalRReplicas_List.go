package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRReplicas_List.json
func ExampleReplicasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListPager("myResourceGroup", "mySignalRService", nil)
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
		// page.ReplicaList = armsignalr.ReplicaList{
		// 	Value: []*armsignalr.Replica{
		// 		{
		// 			Name: to.Ptr("mySignalRService-eastus"),
		// 			Type: to.Ptr("Microsoft.SignalRService/SignalR"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/SignalR/mySignalRService/replicas/mySignalRService-eastus"),
		// 			SystemData: &armsignalr.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armsignalr.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-02-03T04:05:06Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armsignalr.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armsignalr.ReplicaProperties{
		// 				ProvisioningState: to.Ptr(armsignalr.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armsignalr.ResourceSKU{
		// 				Name: to.Ptr("Premium_P1"),
		// 				Capacity: to.Ptr[int32](1),
		// 				Size: to.Ptr("P1"),
		// 				Tier: to.Ptr(armsignalr.SignalRSKUTierPremium),
		// 			},
		// 	}},
		// }
	}
}
