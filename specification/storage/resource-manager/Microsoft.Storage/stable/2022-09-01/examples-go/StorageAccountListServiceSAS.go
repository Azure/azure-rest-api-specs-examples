package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountListServiceSAS.json
func ExampleAccountsClient_ListServiceSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListServiceSAS(ctx, "res7439", "sto1299", armstorage.ServiceSasParameters{
		CanonicalizedResource:  to.Ptr("/blob/sto1299/music"),
		SharedAccessExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T11:32:48.8457197Z"); return t }()),
		Permissions:            to.Ptr(armstorage.PermissionsL),
		Resource:               to.Ptr(armstorage.SignedResourceC),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
