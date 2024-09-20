package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/partnerTransfersGet.json
func ExamplePartnerTransfersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTransfersClient().Get(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "11111111-1111-1111-1111-111111111111", "aabb123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerTransferDetails = armbilling.PartnerTransferDetails{
	// 	Name: to.Ptr("aabb123"),
	// 	Type: to.Ptr("Microsoft.Billing/transfers"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/BillingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111/transfers/aabb123"),
	// 	Properties: &armbilling.PartnerTransferProperties{
	// 		DetailedTransferStatus: []*armbilling.DetailedTransferStatus{
	// 			{
	// 				ProductID: to.Ptr("subscriptionId"),
	// 				ProductName: to.Ptr("Azure subscription 1"),
	// 				ProductType: to.Ptr(armbilling.ProductTypeAzureSubscription),
	// 				SKUDescription: to.Ptr("MS-AZR-0017G"),
	// 				TransferStatus: to.Ptr(armbilling.ProductTransferStatusInProgress),
	// 			},
	// 			{
	// 				ProductID: to.Ptr("reservedInstanceId"),
	// 				ProductName: to.Ptr("Reservation name"),
	// 				ProductType: to.Ptr(armbilling.ProductTypeAzureReservation),
	// 				SKUDescription: to.Ptr("Standard_D2s_v3;VirtualMachines;P1Y"),
	// 				TransferStatus: to.Ptr(armbilling.ProductTransferStatusInProgress),
	// 		}},
	// 		ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-05T17:32:28.000Z"); return t}()),
	// 		InitiatorCustomerType: to.Ptr(armbilling.InitiatorCustomerTypePartner),
	// 		InitiatorEmailID: to.Ptr("xyz@contoso.com"),
	// 		RecipientEmailID: to.Ptr("user@contoso.com"),
	// 		TransferStatus: to.Ptr(armbilling.TransferStatusInProgress),
	// 	},
	// }
}
