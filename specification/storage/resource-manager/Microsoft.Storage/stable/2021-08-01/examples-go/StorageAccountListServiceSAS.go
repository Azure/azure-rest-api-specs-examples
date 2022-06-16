package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountListServiceSAS.json
func ExampleAccountsClient_ListServiceSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.ListServiceSAS(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.ServiceSasParameters{
			CanonicalizedResource:  to.StringPtr("<canonicalized-resource>"),
			SharedAccessExpiryTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T11:32:48.8457197Z"); return t }()),
			Permissions:            armstorage.Permissions("l").ToPtr(),
			Resource:               armstorage.SignedResource("c").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientListServiceSASResult)
}
