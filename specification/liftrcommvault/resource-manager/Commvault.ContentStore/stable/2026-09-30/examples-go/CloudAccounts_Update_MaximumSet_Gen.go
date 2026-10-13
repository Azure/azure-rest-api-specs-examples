package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-09-30/CloudAccounts_Update_MaximumSet_Gen.json
func ExampleCloudAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudAccountsClient().BeginUpdate(ctx, "rgcommvault", "sample-cloudAccountName", armcommvaultcontentstore.CloudAccount{
		Properties: &armcommvaultcontentstore.CloudAccountProperties{
			Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
				SubscriptionID: to.Ptr("bojdg"),
				OfferDetails: &armcommvaultcontentstore.OfferDetails{
					PublisherID: to.Ptr("loowntsnvoynhzto"),
					OfferID:     to.Ptr("ysmwsuakhwvkosz"),
					PlanID:      to.Ptr("iskbkfpr"),
					PlanName:    to.Ptr("pmmlirssfdvdmywddvtl"),
					TermUnit:    to.Ptr("wbeqzbtvq"),
					TermID:      to.Ptr("qoebjvpjc"),
				},
			},
			User: &armcommvaultcontentstore.UserDetails{
				FirstName:    to.Ptr("dudsiomjk"),
				LastName:     to.Ptr("szqupklkgojwozjo"),
				EmailAddress: to.Ptr("user@example.com"),
				Upn:          to.Ptr("wiwwe"),
				PhoneNumber:  to.Ptr("ebszyfnuyzk"),
			},
			Company: &armcommvaultcontentstore.CompanyProfile{
				JobTitle:    to.Ptr("Backup Operator"),
				CompanyName: to.Ptr("Contoso"),
				Website:     to.Ptr("https://www.contoso.com"),
				Street:      to.Ptr("1 Microsoft Way"),
				City:        to.Ptr("Redmond"),
				Country:     to.Ptr("USA"),
				PostalCode:  to.Ptr("98052"),
				State:       to.Ptr("WA"),
			},
		},
		Identity: &armcommvaultcontentstore.ManagedServiceIdentity{
			Type:                   to.Ptr(armcommvaultcontentstore.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armcommvaultcontentstore.UserAssignedIdentity{},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcommvaultcontentstore.CloudAccountsClientUpdateResponse{
	// 	CloudAccount: armcommvaultcontentstore.CloudAccount{
	// 		Properties: &armcommvaultcontentstore.CloudAccountProperties{
	// 			Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("bojdg"),
	// 				SubscriptionStatus: to.Ptr(armcommvaultcontentstore.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				SaasResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Microsoft.SaaS/resources/commvault-saas"),
	// 				OfferDetails: &armcommvaultcontentstore.OfferDetails{
	// 					PublisherID: to.Ptr("loowntsnvoynhzto"),
	// 					OfferID: to.Ptr("ysmwsuakhwvkosz"),
	// 					PlanID: to.Ptr("iskbkfpr"),
	// 					PlanName: to.Ptr("pmmlirssfdvdmywddvtl"),
	// 					TermUnit: to.Ptr("wbeqzbtvq"),
	// 					TermID: to.Ptr("qoebjvpjc"),
	// 				},
	// 			},
	// 			User: &armcommvaultcontentstore.UserDetails{
	// 				FirstName: to.Ptr("dudsiomjk"),
	// 				LastName: to.Ptr("szqupklkgojwozjo"),
	// 				EmailAddress: to.Ptr("user@example.com"),
	// 				Upn: to.Ptr("wiwwe"),
	// 				PhoneNumber: to.Ptr("ebszyfnuyzk"),
	// 			},
	// 			Company: &armcommvaultcontentstore.CompanyProfile{
	// 				JobTitle: to.Ptr("Backup Operator"),
	// 				CompanyName: to.Ptr("Contoso"),
	// 				Website: to.Ptr("https://www.contoso.com"),
	// 				Street: to.Ptr("1 Microsoft Way"),
	// 				City: to.Ptr("Redmond"),
	// 				Country: to.Ptr("USA"),
	// 				PostalCode: to.Ptr("98052"),
	// 				State: to.Ptr("WA"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 			SsoURL: to.Ptr("o"),
	// 		},
	// 		Identity: &armcommvaultcontentstore.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcommvaultcontentstore.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armcommvaultcontentstore.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("sxzmmidsfbba"),
	// 		ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount"),
	// 		Name: to.Ptr("hgu"),
	// 		Type: to.Ptr("wque"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("wg"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.April, 1, 6, 14, 57, 355000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("hbpzxzzwhqfy"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.April, 1, 6, 14, 57, 355000000, time.UTC)),
	// 		},
	// 	},
	// }
}
