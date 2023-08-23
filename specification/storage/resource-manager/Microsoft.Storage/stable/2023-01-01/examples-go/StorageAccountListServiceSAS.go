package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountListServiceSAS.json
func ExampleAccountsClient_ListServiceSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListServiceSAS(ctx, "res7439", "sto1299", armstorage.ServiceSasParameters{
		CanonicalizedResource:  to.Ptr("/blob/sto1299/music"),
		SharedAccessExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T11:32:48.8457197Z"); return t }()),
		Permissions:            to.Ptr(armstorage.PermissionsL),
		Resource:               to.Ptr(armstorage.SignedResourceC),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListServiceSasResponse = armstorage.ListServiceSasResponse{
	// 	ServiceSasToken: to.Ptr("sv=2015-04-05&sr=c&se=2017-05-24T11%3A32%3A48Z&sp=l&sig=PoF8yBUGixsjzwroLmw7vG3VbGz4KB2woZC2D4C2oio%3D"),
	// }
}
