package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/TrustedIdProviders_CreateOrUpdate.json
func ExampleTrustedIDProvidersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTrustedIDProvidersClient().CreateOrUpdate(ctx, "contosorg", "contosoadla", "test_trusted_id_provider_name", armdatalakestore.CreateOrUpdateTrustedIDProviderParameters{
		Properties: &armdatalakestore.CreateOrUpdateTrustedIDProviderProperties{
			IDProvider: to.Ptr("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrustedIDProvider = armdatalakestore.TrustedIDProvider{
	// 	Name: to.Ptr("test_trusted_id_provider_name"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakestore.TrustedIDProviderProperties{
	// 		IDProvider: to.Ptr("https://sts.windows.net/ea9ec534-a3e3-4e45-ad36-3afc5bb291c1"),
	// 	},
	// }
}
