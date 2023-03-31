package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/accounts-subscription-list-all-accounts.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBySubscriptionPager(nil)
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
		// page.MediaServiceCollection = armmediaservices.MediaServiceCollection{
		// 	Value: []*armmediaservices.MediaService{
		// 		{
		// 			Name: to.Ptr("contosotv"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosotv"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmediaservices.MediaServiceProperties{
		// 				MediaServiceID: to.Ptr("6ac94f91-283c-4492-85a7-57976928c17d"),
		// 				PrivateEndpointConnections: []*armmediaservices.PrivateEndpointConnection{
		// 				},
		// 				StorageAccounts: []*armmediaservices.StorageAccount{
		// 					{
		// 						Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contosotvstore"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("contosomovies"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomovies"),
		// 			Location: to.Ptr("South Central US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmediaservices.MediaServiceProperties{
		// 				MediaServiceID: to.Ptr("72681c0f-9dd1-4f1c-95c9-8a8d7d31c4ee"),
		// 				PrivateEndpointConnections: []*armmediaservices.PrivateEndpointConnection{
		// 				},
		// 				StorageAccounts: []*armmediaservices.StorageAccount{
		// 					{
		// 						Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contosomoviesstore"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("fabrikamnews"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/fabrikamnews"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armmediaservices.MediaServiceProperties{
		// 				MediaServiceID: to.Ptr("d96036f9-4e37-491d-8c29-5bc53a29dfcd"),
		// 				PrivateEndpointConnections: []*armmediaservices.PrivateEndpointConnection{
		// 				},
		// 				StorageAccounts: []*armmediaservices.StorageAccount{
		// 					{
		// 						Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/fabrikam/providers/Microsoft.Storage/storageAccounts/fabrikamnewsstore"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
