package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/85cfba195a19120f309bd292c4261aa53a586adb/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/Namespaces/ListBySubscription.json
func ExampleNamespacesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListAllPager(&armnotificationhubs.NamespacesClientListAllOptions{SkipToken: nil,
		Top: nil,
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
		// page.NamespaceListResult = armnotificationhubs.NamespaceListResult{
		// 	Value: []*armnotificationhubs.NamespaceResource{
		// 		{
		// 			Name: to.Ptr("namespace-2"),
		// 			Type: to.Ptr("Microsoft.NotificationHubs/namespaces"),
		// 			ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/namespace-1"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armnotificationhubs.NamespaceProperties{
		// 				Name: to.Ptr("namespace-1"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.010Z"); return t}()),
		// 				Critical: to.Ptr(false),
		// 				Enabled: to.Ptr(true),
		// 				NamespaceType: to.Ptr(armnotificationhubs.NamespaceTypeNotificationHub),
		// 				NetworkACLs: &armnotificationhubs.NetworkACLs{
		// 					IPRules: []*armnotificationhubs.IPRule{
		// 						{
		// 							IPMask: to.Ptr("185.48.100.00/24"),
		// 							Rights: []*armnotificationhubs.AccessRights{
		// 								to.Ptr(armnotificationhubs.AccessRightsManage),
		// 								to.Ptr(armnotificationhubs.AccessRightsSend),
		// 								to.Ptr(armnotificationhubs.AccessRightsListen)},
		// 						}},
		// 						PublicNetworkRule: &armnotificationhubs.PublicInternetAuthorizationRule{
		// 							Rights: []*armnotificationhubs.AccessRights{
		// 								to.Ptr(armnotificationhubs.AccessRightsListen)},
		// 							},
		// 						},
		// 						PrivateEndpointConnections: []*armnotificationhubs.PrivateEndpointConnectionResource{
		// 						},
		// 						ProvisioningState: to.Ptr(armnotificationhubs.OperationProvisioningStateSucceeded),
		// 						PublicNetworkAccess: to.Ptr(armnotificationhubs.PublicNetworkAccessEnabled),
		// 						ServiceBusEndpoint: to.Ptr("https://namespace-1.servicebus.windows.net:443/"),
		// 						Status: to.Ptr(armnotificationhubs.NamespaceStatusCreated),
		// 						SubscriptionID: to.Ptr("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40"),
		// 						UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.030Z"); return t}()),
		// 					},
		// 					SKU: &armnotificationhubs.SKU{
		// 						Name: to.Ptr(armnotificationhubs.SKUNameStandard),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("namespace-2"),
		// 					Type: to.Ptr("Microsoft.NotificationHubs/namespaces"),
		// 					ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/namespace-2"),
		// 					Location: to.Ptr("East US"),
		// 					Tags: map[string]*string{
		// 						"tag1": to.Ptr("value1"),
		// 						"tag2": to.Ptr("value2"),
		// 					},
		// 					Properties: &armnotificationhubs.NamespaceProperties{
		// 						Name: to.Ptr("namespace-2"),
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.010Z"); return t}()),
		// 						Critical: to.Ptr(false),
		// 						Enabled: to.Ptr(true),
		// 						NamespaceType: to.Ptr(armnotificationhubs.NamespaceTypeNotificationHub),
		// 						NetworkACLs: &armnotificationhubs.NetworkACLs{
		// 							IPRules: []*armnotificationhubs.IPRule{
		// 							},
		// 							PublicNetworkRule: &armnotificationhubs.PublicInternetAuthorizationRule{
		// 								Rights: []*armnotificationhubs.AccessRights{
		// 									to.Ptr(armnotificationhubs.AccessRightsManage),
		// 									to.Ptr(armnotificationhubs.AccessRightsListen),
		// 									to.Ptr(armnotificationhubs.AccessRightsSend)},
		// 								},
		// 							},
		// 							PrivateEndpointConnections: []*armnotificationhubs.PrivateEndpointConnectionResource{
		// 							},
		// 							ProvisioningState: to.Ptr(armnotificationhubs.OperationProvisioningStateSucceeded),
		// 							PublicNetworkAccess: to.Ptr(armnotificationhubs.PublicNetworkAccessEnabled),
		// 							ServiceBusEndpoint: to.Ptr("https://namespace-2.servicebus.windows.net:443/"),
		// 							Status: to.Ptr(armnotificationhubs.NamespaceStatusCreated),
		// 							SubscriptionID: to.Ptr("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40"),
		// 							UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.030Z"); return t}()),
		// 						},
		// 						SKU: &armnotificationhubs.SKU{
		// 							Name: to.Ptr(armnotificationhubs.SKUNameStandard),
		// 						},
		// 				}},
		// 			}
	}
}
