package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2026-01-01-preview/Reservations_ListBySubscription_MaximumSet_Gen.json
func ExampleReservationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationsClient().NewListBySubscriptionPager(nil)
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
		// page = armpurestorageblock.ReservationsClientListBySubscriptionResponse{
		// 	ReservationListResult: armpurestorageblock.ReservationListResult{
		// 		Value: []*armpurestorageblock.Reservation{
		// 			{
		// 				Properties: &armpurestorageblock.ReservationPropertiesBaseResourceProperties{
		// 					ReservationInternalID: to.Ptr("res-abc123-xyz789"),
		// 					Marketplace: &armpurestorageblock.MarketplaceDetails{
		// 						SubscriptionID: to.Ptr("Read Storage Poolst-sub-123"),
		// 						SubscriptionStatus: to.Ptr(armpurestorageblock.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 						OfferDetails: &armpurestorageblock.OfferDetails{
		// 							PublisherID: to.Ptr("pure_storage"),
		// 							OfferID: to.Ptr("purestorage-block-offer"),
		// 							PlanID: to.Ptr("standard-plan"),
		// 							PlanName: to.Ptr("Standard Plan"),
		// 							TermUnit: to.Ptr("month"),
		// 							TermID: to.Ptr("12-month-term"),
		// 						},
		// 					},
		// 					User: &armpurestorageblock.UserDetails{
		// 						FirstName: to.Ptr("John"),
		// 						LastName: to.Ptr("Doe"),
		// 						EmailAddress: to.Ptr("john.doe@contoso.com"),
		// 						Upn: to.Ptr("john.doe@contoso.com"),
		// 						PhoneNumber: to.Ptr("+1-425-555-1234"),
		// 						CompanyDetails: &armpurestorageblock.CompanyDetails{
		// 							CompanyName: to.Ptr("Contoso Ltd."),
		// 							Address: &armpurestorageblock.Address{
		// 								AddressLine1: to.Ptr("f"),
		// 								AddressLine2: to.Ptr("Suite 100"),
		// 								City: to.Ptr("Redmond"),
		// 								State: to.Ptr("Washington"),
		// 								Country: to.Ptr("United States"),
		// 								PostalCode: to.Ptr("98052"),
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armpurestorageblock.ProvisioningStateSucceeded),
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("reservation-01"),
		// 				Type: to.Ptr("PureStorage.Block/reservations"),
		// 				SystemData: &armpurestorageblock.SystemData{
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("admin@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-16T07:25:56.721Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/providers/PureStorage.Block/operations?api-version=2026-01-01-preview&$skiptoken=abc123"),
		// 	},
		// }
	}
}
