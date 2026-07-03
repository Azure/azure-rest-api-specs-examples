
/**
 * Samples for SecurityAdminConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationList.json
     */
    /**
     * Sample code: List security admin configurations in a network manager.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listSecurityAdminConfigurationsInANetworkManager(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityAdminConfigurations().list("rg1", "testNetworkManager", null, null,
            com.azure.core.util.Context.NONE);
    }
}
