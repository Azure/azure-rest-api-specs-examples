
/**
 * Samples for NetworkManagerRoutingConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * NetworkManagerRoutingConfigurationList.json
     */
    /**
     * Sample code: List routing configurations in a network manager.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listRoutingConfigurationsInANetworkManager(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagerRoutingConfigurations().list("rg1",
            "testNetworkManager", null, null, com.azure.core.util.Context.NONE);
    }
}
