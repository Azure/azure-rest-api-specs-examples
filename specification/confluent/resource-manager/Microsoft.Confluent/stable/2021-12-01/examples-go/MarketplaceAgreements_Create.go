package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/MarketplaceAgreements_Create.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armconfluent.NewMarketplaceAgreementsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		&armconfluent.MarketplaceAgreementsClientCreateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
