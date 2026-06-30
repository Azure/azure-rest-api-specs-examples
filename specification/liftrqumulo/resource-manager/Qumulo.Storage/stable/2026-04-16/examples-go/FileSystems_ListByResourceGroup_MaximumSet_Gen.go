package armqumulo_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/liftrqumulo/armqumulo/v2"
)

// Generated from example definition: 2026-04-16/FileSystems_ListByResourceGroup_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armqumulo.NewClientFactory("C9CC2D2A-5AA0-4839-A85F-18491F2D244A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgQumulo", nil)
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
		// page = armqumulo.FileSystemsClientListByResourceGroupResponse{
		// 	FileSystemResourceListResult: armqumulo.FileSystemResourceListResult{
		// 		Value: []*armqumulo.FileSystemResource{
		// 			{
		// 				Properties: &armqumulo.FileSystemResourceProperties{
		// 					MarketplaceDetails: &armqumulo.MarketplaceDetails{
		// 						MarketplaceSubscriptionID: to.Ptr("vwjzkiurjihwxrhoicenkbxacokvep"),
		// 						PlanID: to.Ptr("vxnyxa"),
		// 						OfferID: to.Ptr("itiocfnteqyuavgmdtnvwvbpectyr"),
		// 						PublisherID: to.Ptr("zfevjvhjiifwxbazta"),
		// 						TermUnit: to.Ptr("lkbiqoqdyqbua"),
		// 						MarketplaceSubscriptionStatus: to.Ptr(armqumulo.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 					},
		// 					ProvisioningState: to.Ptr(armqumulo.ProvisioningStateAccepted),
		// 					StorageSKU: to.Ptr("myzqnfqelxo"),
		// 					UserDetails: &armqumulo.UserDetails{
		// 					},
		// 					DelegatedSubnetID: to.Ptr("kmjaqsfflkjpke"),
		// 					PerformanceTier: to.Ptr("fjgqmkcvjtygcavpbo"),
		// 					ClusterLoginURL: to.Ptr("uzpvkgxrbgtthyxgavsjr"),
		// 					PrivateIPs: []*string{
		// 						to.Ptr("qrhvbdfbfdgqqe"),
		// 					},
		// 					AvailabilityZone: to.Ptr("luklrtwmngwnaerygykcbwljeso"),
		// 				},
		// 				Identity: &armqumulo.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					Type: to.Ptr(armqumulo.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armqumulo.UserAssignedIdentity{
		// 						"key8111": &armqumulo.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 							ClientID: to.Ptr("33333333-3333-3333-3333-333333333333"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key7800": to.Ptr("izzovtmtunlwpglmglq"),
		// 				},
		// 				Location: to.Ptr("uiuztlexlmxsqcnsdpvzzygmi"),
		// 				ID: to.Ptr("v"),
		// 				Name: to.Ptr("gjaivfviomkvrppi"),
		// 				Type: to.Ptr("njagndjdhsfavdnjqonw"),
		// 				SystemData: &armqumulo.SystemData{
		// 					CreatedBy: to.Ptr("qeoalbtntjkevtpaomizaorvfafj"),
		// 					CreatedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-03T09:55:02.755Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("fdkvrlzayncutzhrhewm"),
		// 					LastModifiedByType: to.Ptr(armqumulo.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-03T09:55:02.755Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/afexan"),
		// 	},
		// }
	}
}
