package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/CloudAccounts_LinkSaaS_MaximumSet_Gen.json
func ExampleCloudAccountsClient_BeginLinkSaaS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudAccountsClient().BeginLinkSaaS(ctx, "rg-commvault", "contoso-cloud-account", armcommvaultcontentstore.SaaSData{
		SaaSResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Microsoft.SaaS/resources/commvault-saas"),
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
	// res = armcommvaultcontentstore.CloudAccountsClientLinkSaaSResponse{
	// 	CloudAccount: armcommvaultcontentstore.CloudAccount{
	// 		Properties: &armcommvaultcontentstore.CloudAccountProperties{
	// 			Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				SubscriptionStatus: to.Ptr(armcommvaultcontentstore.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				SaasResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Microsoft.SaaS/resources/commvault-saas"),
	// 				OfferDetails: &armcommvaultcontentstore.OfferDetails{
	// 					PublisherID: to.Ptr("commvault"),
	// 					OfferID: to.Ptr("commvault-contentstore"),
	// 					PlanID: to.Ptr("commvault-backup-plan-id"),
	// 					PlanName: to.Ptr("Commvault Backup Plan"),
	// 					TermUnit: to.Ptr("P1M"),
	// 					TermID: to.Ptr("term01"),
	// 				},
	// 			},
	// 			User: &armcommvaultcontentstore.UserDetails{
	// 				FirstName: to.Ptr("John"),
	// 				LastName: to.Ptr("Doe"),
	// 				EmailAddress: to.Ptr("john.doe@example.com"),
	// 				Upn: to.Ptr("john.doe@contoso.com"),
	// 				PhoneNumber: to.Ptr("+1-555-0100"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 			SsoURL: to.Ptr("https://cloud.commvault.com/webconsole/sso"),
	// 		},
	// 		Identity: &armcommvaultcontentstore.ManagedServiceIdentity{
	// 			Type: to.Ptr(armcommvaultcontentstore.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armcommvaultcontentstore.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-commvault/providers/Commvault.ContentStore/cloudAccounts/contoso-cloud-account"),
	// 		Name: to.Ptr("contoso-cloud-account"),
	// 		Type: to.Ptr("Commvault.ContentStore/cloudAccounts"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("john.doe@example.com"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T05:18:18.058Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("john.doe@example.com"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-12T05:18:18.059Z"); return t}()),
	// 		},
	// 	},
	// }
}
