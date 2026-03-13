package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps/v2"
)

// Generated from example definition: 2025-10-01-preview/AccountListSAS.json
func ExampleAccountsClient_ListSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("21a9967a-e8a9-4656-a70b-96ff1c4d05a0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListSas(ctx, "myResourceGroup", "myMapsAccount", armmaps.AccountSasParameters{
		Expiry:           to.Ptr("2017-05-24T11:42:03.1567373Z"),
		MaxRatePerSecond: to.Ptr[int32](500),
		PrincipalID:      to.Ptr("e917f87b-324d-4728-98ed-e31d311a7d65"),
		Regions: []*string{
			to.Ptr("eastus"),
		},
		SigningKey: to.Ptr(armmaps.SigningKeyPrimaryKey),
		Start:      to.Ptr("2017-05-24T10:42:03.1567373Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmaps.AccountsClientListSasResponse{
	// 	AccountSasToken: &armmaps.AccountSasToken{
	// 		AccountSasToken: to.Ptr("accountSasToken"),
	// 	},
	// }
}
