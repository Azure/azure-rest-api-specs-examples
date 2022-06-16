package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-04-01/examples/DataCollectionRuleAssociationsDelete.json
func ExampleDataCollectionRuleAssociationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewDataCollectionRuleAssociationsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-uri>",
		"<association-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
