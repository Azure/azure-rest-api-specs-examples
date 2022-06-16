package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleList.json
func ExampleDataMaskingRulesClient_ListBySQLPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewDataMaskingRulesClient("<subscription-id>", cred, nil)
	_, err = client.ListBySQLPool(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<sql-pool-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
