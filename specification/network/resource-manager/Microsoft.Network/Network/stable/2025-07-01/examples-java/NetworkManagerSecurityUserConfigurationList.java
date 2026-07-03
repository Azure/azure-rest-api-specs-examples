
/**
 * Samples for SecurityUserConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserConfigurationList.json
     */
    /**
     * Sample code: List security user configurations in a network manager.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listSecurityUserConfigurationsInANetworkManager(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserConfigurations().list("rg1", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
