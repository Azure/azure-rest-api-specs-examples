package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/OrderGet.json
func ExampleOrdersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrdersClient().Get(ctx, "testedgedevice", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Order = armdataboxedge.Order{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("dataBoxEdgeDevices/orders"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/orders/default"),
	// 	Properties: &armdataboxedge.OrderProperties{
	// 		ContactInformation: &armdataboxedge.ContactDetails{
	// 			CompanyName: to.Ptr("Microsoft"),
	// 			ContactPerson: to.Ptr("John Mcclane"),
	// 			EmailList: []*string{
	// 				to.Ptr("john@microsoft.com")},
	// 				Phone: to.Ptr("(800) 426-9400"),
	// 			},
	// 			CurrentStatus: &armdataboxedge.OrderStatus{
	// 				Comments: to.Ptr(""),
	// 				Status: to.Ptr(armdataboxedge.OrderStateUntracked),
	// 				UpdateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-10T07:59:05.847Z"); return t}()),
	// 			},
	// 			DeliveryTrackingInfo: []*armdataboxedge.TrackingInfo{
	// 			},
	// 			OrderHistory: []*armdataboxedge.OrderStatus{
	// 				{
	// 					Comments: to.Ptr("lorem ipsum"),
	// 					Status: to.Ptr(armdataboxedge.OrderStateUntracked),
	// 					UpdateDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-10T07:59:05.847Z"); return t}()),
	// 			}},
	// 			ReturnTrackingInfo: []*armdataboxedge.TrackingInfo{
	// 			},
	// 			SerialNumber: to.Ptr("UDS123NSDA123"),
	// 			ShippingAddress: &armdataboxedge.Address{
	// 				AddressLine1: to.Ptr("Microsoft Corporation"),
	// 				AddressLine2: to.Ptr("One Microsoft Way"),
	// 				AddressLine3: to.Ptr("Redmond"),
	// 				City: to.Ptr("WA"),
	// 				Country: to.Ptr("USA"),
	// 				PostalCode: to.Ptr("98052"),
	// 				State: to.Ptr("WA"),
	// 			},
	// 		},
	// 	}
}
