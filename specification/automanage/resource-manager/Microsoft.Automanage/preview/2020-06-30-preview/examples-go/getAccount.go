package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/getAccount.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<account-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Account.ID: %s\n", *res.ID)
}
