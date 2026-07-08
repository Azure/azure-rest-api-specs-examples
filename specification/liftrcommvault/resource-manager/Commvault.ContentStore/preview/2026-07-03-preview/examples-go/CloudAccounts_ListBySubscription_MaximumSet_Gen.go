package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/CloudAccounts_ListBySubscription_MaximumSet_Gen.json
func ExampleCloudAccountsClient_NewListBySubscriptionPager_cloudAccountsListBySubscriptionMaximumSetGeneratedByMaximumSetRuleGeneratedByMaximumSetRuleGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudAccountsClient().NewListBySubscriptionPager(nil)
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
		// page = armcommvaultcontentstore.CloudAccountsClientListBySubscriptionResponse{
		// 	CloudAccountListResult: armcommvaultcontentstore.CloudAccountListResult{
		// 		Value: []*armcommvaultcontentstore.CloudAccount{
		// 			{
		// 				Properties: &armcommvaultcontentstore.CloudAccountProperties{
		// 					Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
		// 						SubscriptionID: to.Ptr("tblwyuznrazgchhfczgtlaifwamndt"),
		// 						SubscriptionStatus: to.Ptr(armcommvaultcontentstore.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 						SaasResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Microsoft.SaaS/resources/commvault-saas"),
		// 						OfferDetails: &armcommvaultcontentstore.OfferDetails{
		// 							PublisherID: to.Ptr("npghpdbgiohslbbeihxdwucejb"),
		// 							OfferID: to.Ptr("recofyvhkddgkuvducosjstenmy"),
		// 							PlanID: to.Ptr("pqoyqqavjh"),
		// 							PlanName: to.Ptr("hwcltkdvndwfmmnthzwvocujri"),
		// 							TermUnit: to.Ptr("wzrzqyfzrpqhy"),
		// 							TermID: to.Ptr("avpgkctrkwdmudsz"),
		// 						},
		// 					},
		// 					User: &armcommvaultcontentstore.UserDetails{
		// 						FirstName: to.Ptr("mpiviyooskqkyjqqpgnkderu"),
		// 						LastName: to.Ptr("ppkcvfjylquebr"),
		// 						EmailAddress: to.Ptr("user@example.com"),
		// 						Upn: to.Ptr("frlpmyk"),
		// 						PhoneNumber: to.Ptr("mpunfyfckyzpqxotsmclzk"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
		// 					SsoURL: to.Ptr("o"),
		// 				},
		// 				Identity: &armcommvaultcontentstore.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
		// 					Type: to.Ptr(armcommvaultcontentstore.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armcommvaultcontentstore.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("sxzmmidsfbba"),
		// 				ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/myCloudAccount"),
		// 				Name: to.Ptr("hgu"),
		// 				Type: to.Ptr("wque"),
		// 				SystemData: &armcommvaultcontentstore.SystemData{
		// 					CreatedBy: to.Ptr("wg"),
		// 					CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("hbpzxzzwhqfy"),
		// 					LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-04-01T06:14:57.355Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
