
/**
 * Samples for SecurityUserConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserConfigurationGet.json
     */
    /**
     * Sample code: Get security user configurations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSecurityUserConfigurations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserConfigurations().getWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", com.azure.core.util.Context.NONE);
    }
}
