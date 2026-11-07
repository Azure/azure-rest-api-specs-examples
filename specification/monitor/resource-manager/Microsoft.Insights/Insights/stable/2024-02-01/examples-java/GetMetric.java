
import java.time.Duration;

/**
 * Samples for Metrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetric.json
     */
    /**
     * Sample code: Get Metric for data.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricForData(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listWithResponse(
            "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default",
            "2021-04-20T09:00:00.000Z/2021-04-20T14:00:00.000Z", Duration.parse("PT6H"), "BlobCount,BlobCapacity",
            "average,minimum,maximum", 5, "average asc", "Tier eq '*'", null,
            "Microsoft.Storage/storageAccounts/blobServices", true, false, null, com.azure.core.util.Context.NONE);
    }
}
