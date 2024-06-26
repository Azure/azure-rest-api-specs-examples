package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/UeInfoList.json
func ExampleUeInformationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUeInformationClient().NewListPager("rg1", "TestPacketCoreCP", nil)
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
		// page.UeInfoList = armmobilenetwork.UeInfoList{
		// 	Value: []*armmobilenetwork.UeInfo{
		// 		{
		// 			Name: to.Ptr("001016789123456"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlanes/ues"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/ues/001016789123456"),
		// 			Properties: &armmobilenetwork.UeInfoPropertiesFormat{
		// 				LastReadAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				RatType: to.Ptr(armmobilenetwork.RatTypeFourG),
		// 				UeIPAddresses: []*armmobilenetwork.DnnIPPair{
		// 					{
		// 						Dnn: to.Ptr("Dnn1"),
		// 						UeIPAddress: &armmobilenetwork.UeIPAddress{
		// 							IPV4Addr: to.Ptr("10.0.0.1"),
		// 						},
		// 					},
		// 					{
		// 						Dnn: to.Ptr("Dnn2"),
		// 						UeIPAddress: &armmobilenetwork.UeIPAddress{
		// 							IPV4Addr: to.Ptr("10.0.0.2"),
		// 						},
		// 				}},
		// 				UeState: to.Ptr(armmobilenetwork.UeStateConnected),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("001016789123457"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlanes/ues"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/ues/001016789123457"),
		// 			Properties: &armmobilenetwork.UeInfoPropertiesFormat{
		// 				LastReadAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				RatType: to.Ptr(armmobilenetwork.RatTypeFourG),
		// 				UeIPAddresses: []*armmobilenetwork.DnnIPPair{
		// 					{
		// 						Dnn: to.Ptr("Dnn1"),
		// 						UeIPAddress: &armmobilenetwork.UeIPAddress{
		// 							IPV4Addr: to.Ptr("10.0.0.1"),
		// 						},
		// 					},
		// 					{
		// 						Dnn: to.Ptr("Dnn2"),
		// 						UeIPAddress: &armmobilenetwork.UeIPAddress{
		// 							IPV4Addr: to.Ptr("10.0.0.2"),
		// 						},
		// 				}},
		// 				UeState: to.Ptr(armmobilenetwork.UeStateIdle),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("001016789123458"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/packetCoreControlPlanes/ues"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/ues/001016789123458"),
		// 			Properties: &armmobilenetwork.UeInfoPropertiesFormat{
		// 				LastReadAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				RatType: to.Ptr(armmobilenetwork.RatTypeFourG),
		// 				UeState: to.Ptr(armmobilenetwork.UeStateDetached),
		// 			},
		// 	}},
		// }
	}
}
