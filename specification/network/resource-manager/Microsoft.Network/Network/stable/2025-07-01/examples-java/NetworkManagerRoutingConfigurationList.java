
/**
 * Samples for NetworkManagerRoutingConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerRoutingConfigurationList.json
     */
    /**
     * Sample code: List routing configurations in a network manager.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listRoutingConfigurationsInANetworkManager(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkManagerRoutingConfigurations().list("rg1", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
