
/**
 * Samples for SecurityAdminConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationGet.json
     */
    /**
     * Sample code: Get security admin configurations.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSecurityAdminConfigurations(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityAdminConfigurations().getWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", com.azure.core.util.Context.NONE);
    }
}
