/** Samples for Endpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsList.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsGet.
     *
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsGet(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager
            .endpoints()
            .list(
                "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
                com.azure.core.util.Context.NONE);
    }
}
