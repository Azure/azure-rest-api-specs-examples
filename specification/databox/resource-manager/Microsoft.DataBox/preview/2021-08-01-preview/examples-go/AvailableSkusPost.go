package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/AvailableSkusPost.json
func ExampleServiceClient_ListAvailableSKUsByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewServiceClient("<subscription-id>", cred, nil)
	pager := client.ListAvailableSKUsByResourceGroup("<resource-group-name>",
		"<location>",
		armdatabox.AvailableSKURequest{
			Country:      to.StringPtr("<country>"),
			Location:     to.StringPtr("<location>"),
			TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
		},
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
