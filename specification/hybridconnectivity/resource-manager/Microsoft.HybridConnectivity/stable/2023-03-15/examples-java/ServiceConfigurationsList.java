/** Samples for ServiceConfigurations ListByEndpointResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/ServiceConfigurationsList.json
     */
    /**
     * Sample code: GetClustersExample.
     *
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void getClustersExample(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager
            .serviceConfigurations()
            .listByEndpointResource(
                "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
                "default",
                com.azure.core.util.Context.NONE);
    }
}
