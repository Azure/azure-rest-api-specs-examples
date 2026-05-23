package armnotificationhubs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
)

// Generated from example definition: 2023-10-01-preview/Namespaces/ListByResourceGroup.json
func ExampleNamespacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnotificationhubs.NewClientFactory("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespacesClient().NewListPager("5ktrial", nil)
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
		// page = armnotificationhubs.NamespacesClientListResponse{
		// 	NamespaceListResult: armnotificationhubs.NamespaceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/providers/Microsoft.NotificationHubs/namespaces?api-version=2023-10-01-preview&$count=2&$skipToken=fake-continuation-token-12345"),
		// 		Value: []*armnotificationhubs.NamespaceResource{
		// 			{
		// 				Name: to.Ptr("namespace-2"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/namespace-1"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armnotificationhubs.NamespaceProperties{
		// 					Name: to.Ptr("namespace-1"),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.01+00:00"); return t}()),
		// 					Critical: to.Ptr(false),
		// 					Enabled: to.Ptr(true),
		// 					NamespaceType: to.Ptr(armnotificationhubs.NamespaceTypeNotificationHub),
		// 					NetworkACLs: &armnotificationhubs.NetworkACLs{
		// 						IPRules: []*armnotificationhubs.IPRule{
		// 							{
		// 								IPMask: to.Ptr("185.48.100.00/24"),
		// 								Rights: []*armnotificationhubs.AccessRights{
		// 									to.Ptr(armnotificationhubs.AccessRightsManage),
		// 									to.Ptr(armnotificationhubs.AccessRightsSend),
		// 									to.Ptr(armnotificationhubs.AccessRightsListen),
		// 								},
		// 							},
		// 						},
		// 						PublicNetworkRule: &armnotificationhubs.PublicInternetAuthorizationRule{
		// 							Rights: []*armnotificationhubs.AccessRights{
		// 								to.Ptr(armnotificationhubs.AccessRightsListen),
		// 							},
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armnotificationhubs.PrivateEndpointConnectionResource{
		// 					},
		// 					ProvisioningState: to.Ptr(armnotificationhubs.OperationProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armnotificationhubs.PublicNetworkAccessEnabled),
		// 					ServiceBusEndpoint: to.Ptr("https://namespace-1.servicebus.windows.net:443/"),
		// 					Status: to.Ptr(armnotificationhubs.NamespaceStatusCreated),
		// 					SubscriptionID: to.Ptr("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40"),
		// 					UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.03+00:00"); return t}()),
		// 				},
		// 				SKU: &armnotificationhubs.SKU{
		// 					Name: to.Ptr(armnotificationhubs.SKUNameStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("namespace-2"),
		// 				Type: to.Ptr("Microsoft.NotificationHubs/namespaces"),
		// 				ID: to.Ptr("/subscriptions/29cfa613-cbbc-4512-b1d6-1b3a92c7fa40/resourceGroups/5ktrial/providers/Microsoft.NotificationHubs/namespaces/namespace-2"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armnotificationhubs.NamespaceProperties{
		// 					Name: to.Ptr("namespace-2"),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.01+00:00"); return t}()),
		// 					Critical: to.Ptr(false),
		// 					Enabled: to.Ptr(true),
		// 					NamespaceType: to.Ptr(armnotificationhubs.NamespaceTypeNotificationHub),
		// 					NetworkACLs: &armnotificationhubs.NetworkACLs{
		// 						IPRules: []*armnotificationhubs.IPRule{
		// 						},
		// 						PublicNetworkRule: &armnotificationhubs.PublicInternetAuthorizationRule{
		// 							Rights: []*armnotificationhubs.AccessRights{
		// 								to.Ptr(armnotificationhubs.AccessRightsManage),
		// 								to.Ptr(armnotificationhubs.AccessRightsListen),
		// 								to.Ptr(armnotificationhubs.AccessRightsSend),
		// 							},
		// 						},
		// 					},
		// 					PrivateEndpointConnections: []*armnotificationhubs.PrivateEndpointConnectionResource{
		// 					},
		// 					ProvisioningState: to.Ptr(armnotificationhubs.OperationProvisioningStateSucceeded),
		// 					PublicNetworkAccess: to.Ptr(armnotificationhubs.PublicNetworkAccessEnabled),
		// 					ServiceBusEndpoint: to.Ptr("https://namespace-2.servicebus.windows.net:443/"),
		// 					Status: to.Ptr(armnotificationhubs.NamespaceStatusCreated),
		// 					SubscriptionID: to.Ptr("29cfa613-cbbc-4512-b1d6-1b3a92c7fa40"),
		// 					UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-04-26T06:10:44.03+00:00"); return t}()),
		// 				},
		// 				SKU: &armnotificationhubs.SKU{
		// 					Name: to.Ptr(armnotificationhubs.SKUNameStandard),
		// 				},
		// 				Tags: map[string]*string{
		// 					"tag1": to.Ptr("value1"),
		// 					"tag2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
