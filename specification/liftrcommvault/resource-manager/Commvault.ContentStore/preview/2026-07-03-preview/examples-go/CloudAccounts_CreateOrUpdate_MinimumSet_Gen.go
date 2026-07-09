package armcommvaultcontentstore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/commvaultcontentstore/armcommvaultcontentstore"
)

// Generated from example definition: 2026-07-03-preview/CloudAccounts_CreateOrUpdate_MinimumSet_Gen.json
func ExampleCloudAccountsClient_BeginCreateOrUpdate_cloudAccountsCreateOrUpdateMinimumSetCcaCreateWithCreateOnlyRoleBootstrapFieldsOmitted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommvaultcontentstore.NewClientFactory("65D4E6D7-7063-4C4B-BAC5-13C45474009E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudAccountsClient().BeginCreateOrUpdate(ctx, "rgcommvault", "sample-cloudAccountName", armcommvaultcontentstore.CloudAccount{
		Properties: &armcommvaultcontentstore.CloudAccountProperties{
			Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
				SubscriptionID:     to.Ptr("tblwyuznrazgchhfczgtlaifwamndt"),
				SubscriptionStatus: to.Ptr(armcommvaultcontentstore.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferDetails: &armcommvaultcontentstore.OfferDetails{
					PublisherID: to.Ptr("npghpdbgiohslbbeihxdwucejb"),
					OfferID:     to.Ptr("recofyvhkddgkuvducosjstenmy"),
					PlanID:      to.Ptr("pqoyqqavjh"),
					PlanName:    to.Ptr("hwcltkdvndwfmmnthzwvocujri"),
					TermUnit:    to.Ptr("wzrzqyfzrpqhy"),
					TermID:      to.Ptr("avpgkctrkwdmudsz"),
				},
			},
			User: &armcommvaultcontentstore.UserDetails{
				FirstName:    to.Ptr("John"),
				LastName:     to.Ptr("Doe"),
				EmailAddress: to.Ptr("john.doe@contoso.com"),
				Upn:          to.Ptr("john.doe@contoso.com"),
				PhoneNumber:  to.Ptr("1234567890"),
			},
		},
		Identity: &armcommvaultcontentstore.ManagedServiceIdentity{
			Type:                   to.Ptr(armcommvaultcontentstore.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armcommvaultcontentstore.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("eastus"),
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
	// res = armcommvaultcontentstore.CloudAccountsClientCreateOrUpdateResponse{
	// 	CloudAccount: armcommvaultcontentstore.CloudAccount{
	// 		Properties: &armcommvaultcontentstore.CloudAccountProperties{
	// 			Marketplace: &armcommvaultcontentstore.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("tblwyuznrazgchhfczgtlaifwamndt"),
	// 				SubscriptionStatus: to.Ptr(armcommvaultcontentstore.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armcommvaultcontentstore.OfferDetails{
	// 					PublisherID: to.Ptr("npghpdbgiohslbbeihxdwucejb"),
	// 					OfferID: to.Ptr("recofyvhkddgkuvducosjstenmy"),
	// 					PlanID: to.Ptr("pqoyqqavjh"),
	// 					PlanName: to.Ptr("hwcltkdvndwfmmnthzwvocujri"),
	// 					TermUnit: to.Ptr("wzrzqyfzrpqhy"),
	// 					TermID: to.Ptr("avpgkctrkwdmudsz"),
	// 				},
	// 			},
	// 			User: &armcommvaultcontentstore.UserDetails{
	// 				FirstName: to.Ptr("John"),
	// 				LastName: to.Ptr("Doe"),
	// 				EmailAddress: to.Ptr("john.doe@contoso.com"),
	// 				Upn: to.Ptr("john.doe@contoso.com"),
	// 				PhoneNumber: to.Ptr("1234567890"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcommvaultcontentstore.ResourceProvisioningStateSucceeded),
	// 			SsoURL: to.Ptr("https://cloud.commvault.com/webconsole"),
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
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/65D4E6D7-7063-4C4B-BAC5-13C45474009E/resourceGroups/rgcommvault/providers/Commvault.ContentStore/cloudAccounts/sample-cloudAccountName"),
	// 		Name: to.Ptr("sample-cloudAccountName"),
	// 		Type: to.Ptr("Commvault.ContentStore/cloudAccounts"),
	// 		SystemData: &armcommvaultcontentstore.SystemData{
	// 			CreatedBy: to.Ptr("john.doe@contoso.com"),
	// 			CreatedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("john.doe@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcommvaultcontentstore.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-22T10:17:03.241Z"); return t}()),
	// 		},
	// 	},
	// }
}
