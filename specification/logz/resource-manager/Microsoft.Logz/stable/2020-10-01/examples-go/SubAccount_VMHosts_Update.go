package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccount_VMHosts_Update.json
func ExampleSubAccountClient_ListVMHostUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountClient("<subscription-id>", cred, nil)
	pager := client.ListVMHostUpdate("<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		&armlogz.SubAccountListVMHostUpdateOptions{Body: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("VMResources.ID: %s\n", *v.ID)
		}
	}
}
