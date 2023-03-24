package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSub_ListSkus.json
func ExampleClient_ListSKUs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListSKUs(ctx, "myResourceGroup", "myWebPubSubService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUList = armwebpubsub.SKUList{
	// 	Value: []*armwebpubsub.SKU{
	// 		{
	// 			Capacity: &armwebpubsub.SKUCapacity{
	// 				Default: to.Ptr[int32](1),
	// 				AllowedValues: []*int32{
	// 					to.Ptr[int32](0),
	// 					to.Ptr[int32](1)},
	// 					Maximum: to.Ptr[int32](1),
	// 					Minimum: to.Ptr[int32](0),
	// 					ScaleType: to.Ptr(armwebpubsub.ScaleTypeManual),
	// 				},
	// 				ResourceType: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 				SKU: &armwebpubsub.ResourceSKU{
	// 					Name: to.Ptr("Free_F1"),
	// 					Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierFree),
	// 				},
	// 			},
	// 			{
	// 				Capacity: &armwebpubsub.SKUCapacity{
	// 					Default: to.Ptr[int32](1),
	// 					AllowedValues: []*int32{
	// 						to.Ptr[int32](0),
	// 						to.Ptr[int32](1),
	// 						to.Ptr[int32](2),
	// 						to.Ptr[int32](5),
	// 						to.Ptr[int32](10),
	// 						to.Ptr[int32](20),
	// 						to.Ptr[int32](50),
	// 						to.Ptr[int32](100)},
	// 						Maximum: to.Ptr[int32](100),
	// 						Minimum: to.Ptr[int32](0),
	// 						ScaleType: to.Ptr(armwebpubsub.ScaleTypeAutomatic),
	// 					},
	// 					ResourceType: to.Ptr("Microsoft.SignalRService/WebPubSub"),
	// 					SKU: &armwebpubsub.ResourceSKU{
	// 						Name: to.Ptr("Standard_S1"),
	// 						Tier: to.Ptr(armwebpubsub.WebPubSubSKUTierStandard),
	// 					},
	// 			}},
	// 		}
}
