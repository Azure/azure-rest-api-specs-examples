
/**
 * Samples for AzureTrafficCollectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/AzureTrafficCollectorDelete.json
     */
    /**
     * Sample code: Delete Traffic Collector.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void
        deleteTrafficCollector(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectors().delete("rg1", "atc", com.azure.core.util.Context.NONE);
    }
}
