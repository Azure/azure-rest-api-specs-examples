
/**
 * Samples for AzureTrafficCollectors GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/AzureTrafficCollectorGet.json
     */
    /**
     * Sample code: Get Traffic Collector.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void
        getTrafficCollector(com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectors().getByResourceGroupWithResponse("rg1", "atc", com.azure.core.util.Context.NONE);
    }
}
