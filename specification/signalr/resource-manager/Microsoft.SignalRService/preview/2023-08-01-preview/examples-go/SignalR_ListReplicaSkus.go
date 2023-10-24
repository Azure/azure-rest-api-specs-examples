package armsignalr_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalR_ListReplicaSkus.json
func ExampleClient_ListReplicaSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsignalr.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListReplicaSKUs(ctx, "myResourceGroup", "mySignalRService", "mySignalRService-eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUList = armsignalr.SKUList{
	// 	Value: []*armsignalr.SKU{
	// 		{
	// 			Capacity: &armsignalr.SKUCapacity{
	// 				Default: to.Ptr[int32](1),
	// 				AllowedValues: []*int32{
	// 					to.Ptr[int32](1)},
	// 					Maximum: to.Ptr[int32](1),
	// 					Minimum: to.Ptr[int32](0),
	// 					ScaleType: to.Ptr(armsignalr.ScaleTypeManual),
	// 				},
	// 				ResourceType: to.Ptr("Microsoft.SignalRService/SignalR/replicas"),
	// 				SKU: &armsignalr.ResourceSKU{
	// 					Name: to.Ptr("Free_F1"),
	// 					Tier: to.Ptr(armsignalr.SignalRSKUTierFree),
	// 				},
	// 			},
	// 			{
	// 				Capacity: &armsignalr.SKUCapacity{
	// 					Default: to.Ptr[int32](1),
	// 					AllowedValues: []*int32{
	// 						to.Ptr[int32](1),
	// 						to.Ptr[int32](2),
	// 						to.Ptr[int32](3),
	// 						to.Ptr[int32](4),
	// 						to.Ptr[int32](5),
	// 						to.Ptr[int32](6),
	// 						to.Ptr[int32](7),
	// 						to.Ptr[int32](8),
	// 						to.Ptr[int32](9),
	// 						to.Ptr[int32](10),
	// 						to.Ptr[int32](20),
	// 						to.Ptr[int32](30),
	// 						to.Ptr[int32](40),
	// 						to.Ptr[int32](50),
	// 						to.Ptr[int32](60),
	// 						to.Ptr[int32](70),
	// 						to.Ptr[int32](80),
	// 						to.Ptr[int32](90),
	// 						to.Ptr[int32](100)},
	// 						Maximum: to.Ptr[int32](100),
	// 						Minimum: to.Ptr[int32](0),
	// 						ScaleType: to.Ptr(armsignalr.ScaleTypeManual),
	// 					},
	// 					ResourceType: to.Ptr("Microsoft.SignalRService/SignalR/replicas"),
	// 					SKU: &armsignalr.ResourceSKU{
	// 						Name: to.Ptr("Standard_S1"),
	// 						Tier: to.Ptr(armsignalr.SignalRSKUTierStandard),
	// 					},
	// 				},
	// 				{
	// 					Capacity: &armsignalr.SKUCapacity{
	// 						Default: to.Ptr[int32](1),
	// 						AllowedValues: []*int32{
	// 							to.Ptr[int32](1),
	// 							to.Ptr[int32](2),
	// 							to.Ptr[int32](3),
	// 							to.Ptr[int32](4),
	// 							to.Ptr[int32](5),
	// 							to.Ptr[int32](6),
	// 							to.Ptr[int32](7),
	// 							to.Ptr[int32](8),
	// 							to.Ptr[int32](9),
	// 							to.Ptr[int32](10),
	// 							to.Ptr[int32](20),
	// 							to.Ptr[int32](30),
	// 							to.Ptr[int32](40),
	// 							to.Ptr[int32](50),
	// 							to.Ptr[int32](60),
	// 							to.Ptr[int32](70),
	// 							to.Ptr[int32](80),
	// 							to.Ptr[int32](90),
	// 							to.Ptr[int32](100)},
	// 							Maximum: to.Ptr[int32](100),
	// 							Minimum: to.Ptr[int32](0),
	// 							ScaleType: to.Ptr(armsignalr.ScaleTypeAutomatic),
	// 						},
	// 						ResourceType: to.Ptr("Microsoft.SignalRService/SignalR/replicas"),
	// 						SKU: &armsignalr.ResourceSKU{
	// 							Name: to.Ptr("Premium_P1"),
	// 							Tier: to.Ptr(armsignalr.SignalRSKUTierStandard),
	// 						},
	// 				}},
	// 			}
}
