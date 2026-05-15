package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/transfersInitiate.json
func ExampleTransfersClient_Initiate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransfersClient().Initiate(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "yyyy-yyyy-yyy-yyy", "aabb123", armbilling.InitiateTransferRequest{
		Properties: &armbilling.InitiateTransferProperties{
			RecipientEmailID: to.Ptr("user@contoso.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.TransfersClientInitiateResponse{
	// 	TransferDetails: armbilling.TransferDetails{
	// 		Name: to.Ptr("aabb123"),
	// 		Type: to.Ptr("Microsoft.Billing/transfers"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy/transfers/aabb123"),
	// 		Properties: &armbilling.TransferProperties{
	// 			ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-05T17:32:28Z"); return t}()),
	// 			InitiatorEmailID: to.Ptr("xyz@contoso.com"),
	// 			RecipientEmailID: to.Ptr("user@contoso.com"),
	// 			TransferStatus: to.Ptr(armbilling.TransferStatusPending),
	// 		},
	// 	},
	// }
}
