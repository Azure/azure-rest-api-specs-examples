package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresRegenerateKey.json
func ExampleConfigurationStoresClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappconfiguration.NewConfigurationStoresClient("c80fb759-c965-4c6a-9110-9b2b2d038882", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegenerateKey(ctx,
		"myResourceGroup",
		"contoso",
		armappconfiguration.RegenerateKeyParameters{
			ID: to.Ptr("439AD01B4BE67DB1"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
