package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceById.json
func ExampleInvoicesClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewInvoicesClient("<subscription-id>", cred, nil)
	res, err := client.GetByID(ctx,
		"<invoice-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Invoice.ID: %s\n", *res.ID)
}
