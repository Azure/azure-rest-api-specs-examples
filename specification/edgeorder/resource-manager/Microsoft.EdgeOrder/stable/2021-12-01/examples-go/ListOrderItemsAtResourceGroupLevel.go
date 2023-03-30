package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListOrderItemsAtResourceGroupLevel.json
func ExampleManagementClient_NewListOrderItemsAtResourceGroupLevelPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armedgeorder.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagementClient().NewListOrderItemsAtResourceGroupLevelPager("YourResourceGroupName", &armedgeorder.ManagementClientListOrderItemsAtResourceGroupLevelOptions{Filter: nil,
		Expand:    nil,
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
		// page.OrderItemResourceList = armedgeorder.OrderItemResourceList{
		// 	Value: []*armedgeorder.OrderItemResource{
		// 		{
		// 			Name: to.Ptr("TestOrderItemName1"),
		// 			Type: to.Ptr("Microsoft.EdgeOrder/orderItems"),
		// 			ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armedgeorder.OrderItemProperties{
		// 				AddressDetails: &armedgeorder.AddressDetails{
		// 					ForwardAddress: &armedgeorder.AddressProperties{
		// 						AddressValidationStatus: to.Ptr(armedgeorder.AddressValidationStatusValid),
		// 						ContactDetails: &armedgeorder.ContactDetails{
		// 							ContactName: to.Ptr("XXXX XXXX"),
		// 							EmailList: []*string{
		// 								to.Ptr("xxxx@xxxx.xxx")},
		// 								Phone: to.Ptr("0000000000"),
		// 								PhoneExtension: to.Ptr(""),
		// 							},
		// 							ShippingAddress: &armedgeorder.ShippingAddress{
		// 								AddressType: to.Ptr(armedgeorder.AddressTypeNone),
		// 								City: to.Ptr("San Francisco"),
		// 								CompanyName: to.Ptr("Microsoft"),
		// 								Country: to.Ptr("US"),
		// 								PostalCode: to.Ptr("94107"),
		// 								StateOrProvince: to.Ptr("CA"),
		// 								StreetAddress1: to.Ptr("16 TOWNSEND ST"),
		// 								StreetAddress2: to.Ptr("UNIT 1"),
		// 							},
		// 						},
		// 					},
		// 					OrderID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName1"),
		// 					OrderItemDetails: &armedgeorder.OrderItemDetails{
		// 						CancellationStatus: to.Ptr(armedgeorder.OrderItemCancellationEnumNotCancellable),
		// 						CurrentStage: &armedgeorder.StageDetails{
		// 							StageName: to.Ptr(armedgeorder.StageNameConfirmed),
		// 							StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:05:00.3575338+05:30"); return t}()),
		// 						},
		// 						DeletionStatus: to.Ptr(armedgeorder.ActionStatusEnumNotAllowed),
		// 						ManagementRpDetailsList: []*armedgeorder.ResourceProviderDetails{
		// 							{
		// 								ResourceProviderNamespace: to.Ptr("Microsoft.DataBoxEdge"),
		// 						}},
		// 						NotificationEmailList: []*string{
		// 						},
		// 						OrderItemStageHistory: []*armedgeorder.StageDetails{
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:55:54.3427968+05:30"); return t}()),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameConfirmed),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:05:00.3575338+05:30"); return t}()),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameReadyToShip),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameShipped),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameDelivered),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 							},
		// 							{
		// 								StageName: to.Ptr(armedgeorder.StageNameInUse),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 						}},
		// 						OrderItemType: to.Ptr(armedgeorder.OrderItemTypePurchase),
		// 						Preferences: &armedgeorder.Preferences{
		// 							TransportPreferences: &armedgeorder.TransportPreferences{
		// 								PreferredShipmentType: to.Ptr(armedgeorder.TransportShipmentTypesMicrosoftManaged),
		// 							},
		// 						},
		// 						ProductDetails: &armedgeorder.ProductDetails{
		// 							Count: to.Ptr[int32](0),
		// 							DisplayInfo: &armedgeorder.DisplayInfo{
		// 								ConfigurationDisplayName: to.Ptr("Azure Stack Edge Pro - 1 GPU"),
		// 								ProductFamilyDisplayName: to.Ptr("Azure Stack Edge"),
		// 							},
		// 							HierarchyInformation: &armedgeorder.HierarchyInformation{
		// 								ConfigurationName: to.Ptr("edgep_base"),
		// 								ProductFamilyName: to.Ptr("azurestackedge"),
		// 								ProductLineName: to.Ptr("azurestackedge"),
		// 								ProductName: to.Ptr("azurestackedgegpu"),
		// 							},
		// 							ProductDoubleEncryptionStatus: to.Ptr(armedgeorder.DoubleEncryptionStatusDisabled),
		// 						},
		// 						ReturnStatus: to.Ptr(armedgeorder.OrderItemReturnEnumNotReturnable),
		// 					},
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:55:10.2820482+05:30"); return t}()),
		// 				},
		// 				SystemData: &armedgeorder.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("TestOrderItemName2"),
		// 				Type: to.Ptr("Microsoft.EdgeOrder/orderItems"),
		// 				ID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/orderItems/TestOrderItemName2"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armedgeorder.OrderItemProperties{
		// 					AddressDetails: &armedgeorder.AddressDetails{
		// 						ForwardAddress: &armedgeorder.AddressProperties{
		// 							AddressValidationStatus: to.Ptr(armedgeorder.AddressValidationStatusValid),
		// 							ContactDetails: &armedgeorder.ContactDetails{
		// 								ContactName: to.Ptr("XXXX XXXX"),
		// 								EmailList: []*string{
		// 									to.Ptr("xxxx@xxxx.xxx")},
		// 									Phone: to.Ptr("0000000000"),
		// 									PhoneExtension: to.Ptr(""),
		// 								},
		// 								ShippingAddress: &armedgeorder.ShippingAddress{
		// 									AddressType: to.Ptr(armedgeorder.AddressTypeNone),
		// 									City: to.Ptr("San Francisco"),
		// 									CompanyName: to.Ptr("Microsoft"),
		// 									Country: to.Ptr("US"),
		// 									PostalCode: to.Ptr("94107"),
		// 									StateOrProvince: to.Ptr("CA"),
		// 									StreetAddress1: to.Ptr("16 TOWNSEND ST"),
		// 									StreetAddress2: to.Ptr("UNIT 1"),
		// 								},
		// 							},
		// 						},
		// 						OrderID: to.Ptr("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.EdgeOrder/locations/eastus/orders/TestOrderName2"),
		// 						OrderItemDetails: &armedgeorder.OrderItemDetails{
		// 							CancellationStatus: to.Ptr(armedgeorder.OrderItemCancellationEnumNotCancellable),
		// 							CurrentStage: &armedgeorder.StageDetails{
		// 								StageName: to.Ptr(armedgeorder.StageNameConfirmed),
		// 								StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 								StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:07:29.9896685+05:30"); return t}()),
		// 							},
		// 							DeletionStatus: to.Ptr(armedgeorder.ActionStatusEnumNotAllowed),
		// 							ManagementRpDetailsList: []*armedgeorder.ResourceProviderDetails{
		// 								{
		// 									ResourceProviderNamespace: to.Ptr("Microsoft.DataBoxEdge"),
		// 							}},
		// 							NotificationEmailList: []*string{
		// 							},
		// 							OrderItemStageHistory: []*armedgeorder.StageDetails{
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNamePlaced),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:59:04.9701334+05:30"); return t}()),
		// 								},
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNameConfirmed),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusSucceeded),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T11:07:29.9896685+05:30"); return t}()),
		// 								},
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNameReadyToShip),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 								},
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNameShipped),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 								},
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNameDelivered),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 								},
		// 								{
		// 									StageName: to.Ptr(armedgeorder.StageNameInUse),
		// 									StageStatus: to.Ptr(armedgeorder.StageStatusNone),
		// 							}},
		// 							OrderItemType: to.Ptr(armedgeorder.OrderItemTypePurchase),
		// 							Preferences: &armedgeorder.Preferences{
		// 								TransportPreferences: &armedgeorder.TransportPreferences{
		// 									PreferredShipmentType: to.Ptr(armedgeorder.TransportShipmentTypesMicrosoftManaged),
		// 								},
		// 							},
		// 							ProductDetails: &armedgeorder.ProductDetails{
		// 								Count: to.Ptr[int32](0),
		// 								DisplayInfo: &armedgeorder.DisplayInfo{
		// 									ConfigurationDisplayName: to.Ptr("Azure Stack Edge Pro - 1 GPU"),
		// 									ProductFamilyDisplayName: to.Ptr("Azure Stack Edge"),
		// 								},
		// 								HierarchyInformation: &armedgeorder.HierarchyInformation{
		// 									ConfigurationName: to.Ptr("edgep_base"),
		// 									ProductFamilyName: to.Ptr("azurestackedge"),
		// 									ProductLineName: to.Ptr("azurestackedge"),
		// 									ProductName: to.Ptr("azurestackedgegpu"),
		// 								},
		// 								ProductDoubleEncryptionStatus: to.Ptr(armedgeorder.DoubleEncryptionStatusDisabled),
		// 							},
		// 							ReturnStatus: to.Ptr(armedgeorder.OrderItemReturnEnumNotReturnable),
		// 						},
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T10:58:27.5824859+05:30"); return t}()),
		// 					},
		// 					SystemData: &armedgeorder.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T05:30:00+05:30"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}
