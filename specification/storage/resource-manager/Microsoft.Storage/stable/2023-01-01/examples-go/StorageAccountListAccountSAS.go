package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountListAccountSAS.json
func ExampleAccountsClient_ListAccountSAS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListAccountSAS(ctx, "res7985", "sto8588", armstorage.AccountSasParameters{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListAccountSasResponse = armstorage.ListAccountSasResponse{
	// 	AccountSasToken: to.Ptr("sv=2015-04-05&ss=b&srt=s&sp=r&st=2017-05-24T10%3A42%3A03Z&se=2017-05-24T11%3A42%3A03Z&spr=https,http&sig=Z0I%2BEpM%2BPPlTC8ApfUf%2BcffO2aahMgZim3U0iArqsS0%3D"),
	// }
}
