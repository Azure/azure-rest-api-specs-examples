package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2018-06-01/examples/Pricings/GetPricingByName_example.json
func ExamplePricingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewPricingsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<pricing-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PricingsClientGetResult)
}
