package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-sas-urls.json
func ExampleAssetsClient_ListContainerSas() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetsClient().ListContainerSas(ctx, "contoso", "contosomedia", "ClimbingMountBaker", armmediaservices.ListContainerSasInput{
		ExpiryTime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T10:00:00.007Z"); return t }()),
		Permissions: to.Ptr(armmediaservices.AssetContainerPermissionReadWrite),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssetContainerSas = armmediaservices.AssetContainerSas{
	// 	AssetContainerSasUrls: []*string{
	// 		to.Ptr("https://storage0.blob.core.windows.net/asset-89af1750-e681-4fbe-8c4c-9a5567867a6b?sr=b&sig=<secret-value-from-key-1>&se=2018-01-01T10:00:00Z&sp=lrw"),
	// 		to.Ptr("https://storage0.blob.core.windows.net/asset-89af1750-e681-4fbe-8c4c-9a5567867a6b?sr=b&sig=<secret-value-from-key-2>&se=2018-01-01T10:00:00Z&sp=lrw")},
	// 	}
}
