
import com.azure.resourcemanager.monitor.models.ResultType;
import java.time.Duration;

/**
 * Samples for Metrics List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricMetadata.json
     */
    /**
     * Sample code: Get Metric for metadata.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMetricForMetadata(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetrics().listWithResponse(
            "subscriptions/b324c52b-4073-4807-93af-e07d289c093e/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/larryshoebox/blobServices/default",
            "2017-04-14T02:20:00Z/2017-04-14T04:20:00Z", Duration.parse("PT1M"), null, "Average,count", 3,
            "Average asc", "BlobType eq '*'", ResultType.fromString("metadata"),
            "Microsoft.Storage/storageAccounts/blobServices", com.azure.core.util.Context.NONE);
    }
}
