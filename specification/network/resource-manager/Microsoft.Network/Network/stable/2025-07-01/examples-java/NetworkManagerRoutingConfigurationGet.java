
/**
 * Samples for NetworkManagerRoutingConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingConfigurationGet.json
     */
    /**
     * Sample code: Get routing configurations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getRoutingConfigurations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagerRoutingConfigurations().getWithResponse("rg1", "testNetworkManager",
            "myTestRoutingConfig", com.azure.core.util.Context.NONE);
    }
}
