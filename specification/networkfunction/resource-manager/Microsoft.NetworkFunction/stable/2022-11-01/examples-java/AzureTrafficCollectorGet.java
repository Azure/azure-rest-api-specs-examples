import com.azure.core.util.Context;

/** Samples for AzureTrafficCollectors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/AzureTrafficCollectorGet.json
     */
    /**
     * Sample code: Get Traffic Collector.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void getTrafficCollector(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectors().getByResourceGroupWithResponse("rg1", "atc", Context.NONE);
    }
}
