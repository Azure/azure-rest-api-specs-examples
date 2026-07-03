
/**
 * Samples for SecurityAdminConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager security admin configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteNetworkManagerSecurityAdminConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityAdminConfigurations().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", false, com.azure.core.util.Context.NONE);
    }
}
