package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/RegenerateKey.json
func ExampleAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmaps.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().RegenerateKeys(ctx, "myResourceGroup", "myMapsAccount", armmaps.KeySpecification{
		KeyType: to.Ptr(armmaps.KeyTypePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountKeys = armmaps.AccountKeys{
	// 	PrimaryKey: to.Ptr("<primaryKey>"),
	// 	PrimaryKeyLastUpdated: to.Ptr("2021-07-02T01:01:01.1075056Z"),
	// 	SecondaryKey: to.Ptr("<secondaryKey>"),
	// 	SecondaryKeyLastUpdated: to.Ptr("2021-07-02T01:01:01.1075056Z"),
	// }
}
