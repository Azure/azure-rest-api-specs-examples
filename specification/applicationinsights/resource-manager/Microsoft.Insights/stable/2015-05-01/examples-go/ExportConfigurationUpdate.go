package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationUpdate.json
func ExampleExportConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewExportConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"my-component",
		"uGOoki0jQsyEs3IdQ83Q4QsNr4=",
		armapplicationinsights.ComponentExportRequest{
			DestinationAccountID:             to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob"),
			DestinationAddress:               to.Ptr("https://mystorageblob.blob.core.windows.net/fchentest?sv=2015-04-05&sr=c&sig=token"),
			DestinationStorageLocationID:     to.Ptr("eastus"),
			DestinationStorageSubscriptionID: to.Ptr("subid"),
			DestinationType:                  to.Ptr("Blob"),
			IsEnabled:                        to.Ptr("true"),
			NotificationQueueEnabled:         to.Ptr("false"),
			NotificationQueueURI:             to.Ptr(""),
			RecordTypes:                      to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
