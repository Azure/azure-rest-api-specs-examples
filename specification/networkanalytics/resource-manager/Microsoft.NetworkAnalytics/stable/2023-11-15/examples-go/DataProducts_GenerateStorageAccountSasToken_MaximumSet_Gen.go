package armnetworkanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/networkanalytics/resource-manager/Microsoft.NetworkAnalytics/stable/2023-11-15/examples/DataProducts_GenerateStorageAccountSasToken_MaximumSet_Gen.json
func ExampleDataProductsClient_GenerateStorageAccountSasToken_dataProductsGenerateStorageAccountSasTokenMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataProductsClient().GenerateStorageAccountSasToken(ctx, "aoiresourceGroupName", "dataproduct01", armnetworkanalytics.AccountSas{
		ExpiryTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T05:34:58.151Z"); return t }()),
		IPAddress:       to.Ptr("1.1.1.1"),
		StartTimeStamp:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-24T05:34:58.151Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountSasToken = armnetworkanalytics.AccountSasToken{
	// 	StorageAccountSasToken: to.Ptr("storageAccountSasToken"),
	// }
}
