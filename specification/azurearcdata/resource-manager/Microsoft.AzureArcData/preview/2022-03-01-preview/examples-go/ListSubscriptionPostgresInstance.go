package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListSubscriptionPostgresInstance.json
func ExamplePostgresInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPostgresInstancesClient().NewListPager(nil)
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
		// page.PostgresInstanceListResult = armazurearcdata.PostgresInstanceListResult{
		// 	Value: []*armazurearcdata.PostgresInstance{
		// 		{
		// 			Name: to.Ptr("testpostgresInstances1"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/PostgresInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/PostgresInstances/testpostgresInstances1"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.PostgresInstanceProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 			},
		// 			SKU: &armazurearcdata.PostgresInstanceSKU{
		// 				Name: to.Ptr("default"),
		// 				Dev: to.Ptr(true),
		// 				Tier: to.Ptr("Hyperscale"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testpostgresInstances2"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/PostgresInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/PostgresInstances/testpostgresInstances2"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.PostgresInstanceProperties{
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 			},
		// 			SKU: &armazurearcdata.PostgresInstanceSKU{
		// 				Name: to.Ptr("default"),
		// 				Dev: to.Ptr(true),
		// 				Tier: to.Ptr("Hyperscale"),
		// 			},
		// 	}},
		// }
	}
}
