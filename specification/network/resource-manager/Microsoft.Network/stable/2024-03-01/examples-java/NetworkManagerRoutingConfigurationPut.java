
import com.azure.resourcemanager.network.fluent.models.NetworkManagerRoutingConfigurationInner;

/**
 * Samples for NetworkManagerRoutingConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/
     * NetworkManagerRoutingConfigurationPut.json
     */
    /**
     * Sample code: Create network manager routing configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkManagerRoutingConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkManagerRoutingConfigurations().createOrUpdateWithResponse(
            "rg1", "testNetworkManager", "myTestRoutingConfig",
            new NetworkManagerRoutingConfigurationInner().withDescription("A sample policy"),
            com.azure.core.util.Context.NONE);
    }
}
