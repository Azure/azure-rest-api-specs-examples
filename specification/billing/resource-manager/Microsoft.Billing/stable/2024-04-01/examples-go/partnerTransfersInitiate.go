package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/partnerTransfersInitiate.json
func ExamplePartnerTransfersClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnerTransfersClient().Initiate(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "11111111-1111-1111-1111-111111111111", "aabb123", armbilling.PartnerInitiateTransferRequest{
		Properties: &armbilling.PartnerInitiateTransferProperties{
			RecipientEmailID: to.Ptr("user@contoso.com"),
		},
	}, nil)
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
	// 		ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-05T17:32:28.000Z"); return t}()),
	// 		InitiatorCustomerType: to.Ptr(armbilling.InitiatorCustomerTypePartner),
	// 		InitiatorEmailID: to.Ptr("xyz@contoso.com"),
	// 		RecipientEmailID: to.Ptr("user@contoso.com"),
	// 		TransferStatus: to.Ptr(armbilling.TransferStatusPending),
	// 	},
	// }
}
