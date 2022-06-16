package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/MarketplaceAgreements_Create.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconfluent.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		&armconfluent.MarketplaceAgreementsClientCreateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientCreateResult)
}
