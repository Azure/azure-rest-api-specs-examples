package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AddressInvalid.json
func ExampleAddressClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewAddressClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Validate(ctx,
		armbilling.AddressDetails{
			AddressLine1: to.Ptr("1 Test"),
			City:         to.Ptr("bellevue"),
			Country:      to.Ptr("us"),
			PostalCode:   to.Ptr("12345"),
			Region:       to.Ptr("wa"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
