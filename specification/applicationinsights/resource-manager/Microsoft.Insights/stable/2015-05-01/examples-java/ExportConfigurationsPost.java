
import com.azure.resourcemanager.applicationinsights.models.ApplicationInsightsComponentExportRequest;

/**
 * Samples for ExportConfigurations Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/
     * ExportConfigurationsPost.json
     */
    /**
     * Sample code: ExportConfigurationPost.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void
        exportConfigurationPost(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.exportConfigurations().createWithResponse("my-resource-group", "my-component",
            new ApplicationInsightsComponentExportRequest().withRecordTypes(
                "Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability")
                .withDestinationType("Blob")
                .withDestinationAddress(
                    "https://mystorageblob.blob.core.windows.net/testexport?sv=2015-04-05&sr=c&sig=token")
                .withIsEnabled("true").withNotificationQueueEnabled("false").withNotificationQueueUri("")
                .withDestinationStorageSubscriptionId("subid").withDestinationStorageLocationId("eastus")
                .withDestinationAccountId(
                    "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob"),
            com.azure.core.util.Context.NONE);
    }
}
