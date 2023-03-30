package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/792db17291c758b2bfdbbc0d35d0e2f5b5a1bd05/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuListCredential.json
func ExampleVendorSKUsClient_ListCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVendorSKUsClient().ListCredential(ctx, "TestVendor", "TestSku", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SKUCredential = armhybridnetwork.SKUCredential{
	// 	AcrServerURL: to.Ptr("test.azurecr.io"),
	// 	AcrToken: to.Ptr("dummytoken"),
	// 	Expiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-06-07T18:01:41.334+00:00"); return t}()),
	// 	Repositories: []*string{
	// 		to.Ptr("")},
	// 		Username: to.Ptr("dummyuser"),
	// 	}
}
