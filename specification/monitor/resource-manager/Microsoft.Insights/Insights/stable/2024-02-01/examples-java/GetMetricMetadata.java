
/**
 * Samples for Metrics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01/GetMetricMetadata.json
     */
    /**
     * Sample code: Get Metric for metadata.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getMetricForMetadata(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getMetrics().listWithResponse(
            "subscriptions/1f3fa6d2-851c-4a91-9087-1a050f3a9c38/resourceGroups/todking/providers/Microsoft.Storage/storageAccounts/tkfileserv/blobServices/default",
            "2017-04-14T02:20:00Z/2017-04-14T04:20:00Z", null, null, null, null, null, "Tier eq '*'", null,
            "Microsoft.Storage/storageAccounts/blobServices", null, null, null, com.azure.core.util.Context.NONE);
    }
}
