
/**
 * Samples for SecurityUserConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkManagerSecurityUserConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager security user configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteNetworkManagerSecurityUserConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserConfigurations().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", false, com.azure.core.util.Context.NONE);
    }
}
