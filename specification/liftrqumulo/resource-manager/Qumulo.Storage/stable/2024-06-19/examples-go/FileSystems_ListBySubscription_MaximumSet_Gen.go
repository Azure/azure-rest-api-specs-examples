package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72de08114673a547de8a017c85ed89a2017a86f7/specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_ListBySubscription_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListBySubscriptionPager_fileSystemsListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListBySubscriptionPager(nil)
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
		// page.FileSystemResourceListResult = armqumulo.FileSystemResourceListResult{
		// 	Value: []*armqumulo.FileSystemResource{
		// 		{
		// 			Name: to.Ptr("stftolw"),
		// 			Type: to.Ptr("wj"),
		// 			ID: to.Ptr("rfta"),
		// 			SystemData: &armqumulo.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T08:11:54.895Z"); return t}()),
		// 				CreatedBy: to.Ptr("usnkckwkizihezb"),
		// 				CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T08:11:54.895Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("yjsiqdgtsmycxlncjceemlucn"),
		// 				LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("pnb"),
		// 			Tags: map[string]*string{
		// 				"key7090": to.Ptr("rurrdiaqp"),
		// 			},
		// 			Identity: &armqumulo.ManagedServiceIdentity{
		// 				Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
		// 					"key7679": &armqumulo.UserAssignedIdentity{
		// 						ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 						PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armqumulo.FileSystemResourceProperties{
		// 				AvailabilityZone: to.Ptr("eqdvbdiuwmhhzqzmksmwllpddqquwt"),
		// 				ClusterLoginURL: to.Ptr("ykaynsjvhihdthkkvvodjrgc"),
		// 				DelegatedSubnetID: to.Ptr("jykmxrf"),
		// 				MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 					MarketplaceSubscriptionID: to.Ptr("xaqtkloiyovmexqhn"),
		// 					MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					OfferID: to.Ptr("s"),
		// 					PlanID: to.Ptr("fwtpz"),
		// 					PublisherID: to.Ptr("czxcfrwodazyaft"),
		// 					TermUnit: to.Ptr("cfwwczmygsimcyvoclcw"),
		// 				},
		// 				PrivateIPs: []*string{
		// 					to.Ptr("gzken")},
		// 					ProvisioningState: to.Ptr(armqumulo.ProvisioningStateSucceeded),
		// 					StorageSKU: to.Ptr("yhyzby"),
		// 					UserDetails: &armqumulo.UserDetails{
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
