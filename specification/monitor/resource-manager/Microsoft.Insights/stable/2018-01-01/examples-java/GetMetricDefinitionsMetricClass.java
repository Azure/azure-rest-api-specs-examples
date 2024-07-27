
/**
 * Samples for MetricDefinitions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/
     * GetMetricDefinitionsMetricClass.json
     */
    /**
     * Sample code: Get StorageCache Metric Definitions with metric class.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getStorageCacheMetricDefinitionsWithMetricClass(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricDefinitions().list(
            "subscriptions/46841c0e-69c8-4b17-af46-6626ecb15fc2/resourceGroups/adgarntptestrg/providers/Microsoft.StorageCache/caches/adgarntptestcache",
            "microsoft.storagecache/caches", com.azure.core.util.Context.NONE);
    }
}
