
/**
 * Samples for NetworkManagerRoutingConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager routing configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteNetworkManagerRoutingConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagerRoutingConfigurations().delete("rg1", "testNetworkManager",
            "myTestRoutingConfig", null, com.azure.core.util.Context.NONE);
    }
}
