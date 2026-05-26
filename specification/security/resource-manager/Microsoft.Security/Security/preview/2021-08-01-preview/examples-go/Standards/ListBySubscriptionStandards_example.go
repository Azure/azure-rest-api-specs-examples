package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2021-08-01-preview/Standards/ListBySubscriptionStandards_example.json
func ExampleStandardsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStandardsClient().NewListBySubscriptionPager(nil)
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
		// page = armsecurity.StandardsClientListBySubscriptionResponse{
		// 	ArmSecurityStandardList: armsecurity.ArmSecurityStandardList{
		// 		Value: []*armsecurity.Standard{
		// 			{
		// 				Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Type: to.Ptr("Microsoft.Security/standards"),
		// 				Etag: to.Ptr("etag value"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myResourceGroup/providers/Microsoft.Security/standards/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armsecurity.StandardProperties{
		// 					Description: to.Ptr("Controls and security data implementing security recommendations defined in Azure Security"),
		// 					Category: to.Ptr("SecurityCenter"),
		// 					Components: []*armsecurity.StandardComponentProperties{
		// 						{
		// 							Key: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
		// 						},
		// 						{
		// 							Key: to.Ptr("dbd0cb49-b563-45e7-9724-889e799fa648"),
		// 						},
		// 					},
		// 					DisplayName: to.Ptr("Cross cloud standard 1"),
		// 					StandardType: to.Ptr("Custom"),
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("155afdf9-d239-4a5c-847f-89da613e7143"),
		// 				Type: to.Ptr("Microsoft.Security/standards"),
		// 				Etag: to.Ptr("etag value"),
		// 				ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myResourceGroup/providers/Microsoft.Security/standards/155afdf9-d239-4a5c-847f-89da613e7143"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armsecurity.StandardProperties{
		// 					Description: to.Ptr("Controls and security data implementing security recommendations defined in Azure Security"),
		// 					Category: to.Ptr("SecurityCenter"),
		// 					Components: []*armsecurity.StandardComponentProperties{
		// 						{
		// 							Key: to.Ptr("1195afff-c881-495e-9bc5-1486211ae03f"),
		// 						},
		// 						{
		// 							Key: to.Ptr("dbd0cb49-b563-45e7-9724-889e799fa648"),
		// 						},
		// 					},
		// 					DisplayName: to.Ptr("Cross cloud standard 2"),
		// 					StandardType: to.Ptr("Custom"),
		// 					SupportedClouds: []*armsecurity.StandardSupportedClouds{
		// 						to.Ptr(armsecurity.StandardSupportedCloudsGCP),
		// 					},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
