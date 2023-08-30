/** Samples for Endpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsGetCustom.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsGetCustom.
     *
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsGetCustom(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager
            .endpoints()
            .getWithResponse(
                "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
                "custom",
                com.azure.core.util.Context.NONE);
    }
}
