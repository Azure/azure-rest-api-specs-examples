package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/AccountListSAS.json
func ExampleAccountsClient_ListSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListSas(ctx, "myResourceGroup", "myMapsAccount", armmaps.AccountSasParameters{
		Expiry:           to.Ptr("2017-05-24T11:42:03.1567373Z"),
		MaxRatePerSecond: to.Ptr[int32](500),
		PrincipalID:      to.Ptr("e917f87b-324d-4728-98ed-e31d311a7d65"),
		Regions: []*string{
			to.Ptr("eastus")},
		SigningKey: to.Ptr(armmaps.SigningKeyPrimaryKey),
		Start:      to.Ptr("2017-05-24T10:42:03.1567373Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountSasToken = armmaps.AccountSasToken{
	// 	AccountSasToken: to.Ptr("accountSasToken"),
	// }
}
