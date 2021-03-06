package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolAuditingSettingsList.json
func ExampleSQLPoolBlobAuditingPoliciesClient_ListBySQLPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewSQLPoolBlobAuditingPoliciesClient("<subscription-id>", cred, nil)
	pager := client.ListBySQLPool("<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("SQLPoolBlobAuditingPolicy.ID: %s\n", *v.ID)
		}
	}
}
