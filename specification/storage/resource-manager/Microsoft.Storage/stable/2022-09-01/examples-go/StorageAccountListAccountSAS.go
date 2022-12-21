package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountListAccountSAS.json
func ExampleAccountsClient_ListAccountSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewAccountsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListAccountSAS(ctx, "res7985", "sto8588", armstorage.AccountSasParameters{
		KeyToSign:              to.Ptr("key1"),
		SharedAccessExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T11:42:03.1567373Z"); return t }()),
		Permissions:            to.Ptr(armstorage.PermissionsR),
		Protocols:              to.Ptr(armstorage.HTTPProtocolHTTPSHTTP),
		ResourceTypes:          to.Ptr(armstorage.SignedResourceTypesS),
		Services:               to.Ptr(armstorage.ServicesB),
		SharedAccessStartTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T10:42:03.1567373Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
