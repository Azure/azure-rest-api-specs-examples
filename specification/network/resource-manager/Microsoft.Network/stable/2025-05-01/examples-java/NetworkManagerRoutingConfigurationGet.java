
/**
 * Samples for NetworkManagerRoutingConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerRoutingConfigurationGet.json
     */
    /**
     * Sample code: Get routing configurations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoutingConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagerRoutingConfigurations().getWithResponse("rg1",
            "testNetworkManager", "myTestRoutingConfig", com.azure.core.util.Context.NONE);
    }
}
