package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountListAccountSAS.json
func ExampleAccountsClient_ListAccountSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.ListAccountSAS(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.AccountSasParameters{
			KeyToSign:              to.StringPtr("<key-to-sign>"),
			SharedAccessExpiryTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T11:42:03.1567373Z"); return t }()),
			Permissions:            armstorage.Permissions("r").ToPtr(),
			Protocols:              armstorage.HTTPProtocolHTTPSHTTP.ToPtr(),
			ResourceTypes:          armstorage.SignedResourceTypes("s").ToPtr(),
			Services:               armstorage.Services("b").ToPtr(),
			SharedAccessStartTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T10:42:03.1567373Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientListAccountSASResult)
}
