
/**
 * Samples for ServiceConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/ServiceConfigurationsGetWAC.json
     */
    /**
     * Sample code: HybridConnectivityEndpointsServiceconfigurationsGetWAC.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityEndpointsServiceconfigurationsGetWAC(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.serviceConfigurations().getWithResponse(
            "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default",
            "default", "WAC", com.azure.core.util.Context.NONE);
    }
}
