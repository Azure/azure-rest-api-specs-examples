
/**
 * Samples for NetworkManagerRoutingConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerRoutingConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager routing configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkManagerRoutingConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagerRoutingConfigurations().delete("rg1",
            "testNetworkManager", "myTestRoutingConfig", null, com.azure.core.util.Context.NONE);
    }
}
