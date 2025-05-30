
/**
 * Samples for Endpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/
     * EndpointsDeleteDefault.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsDeleteDefault.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsDeleteDefault(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.endpoints().deleteByResourceGroupWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine",
            "default", com.azure.core.util.Context.NONE);
    }
}
