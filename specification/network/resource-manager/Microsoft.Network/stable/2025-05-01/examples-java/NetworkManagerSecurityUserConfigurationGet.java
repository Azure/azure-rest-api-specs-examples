
/**
 * Samples for SecurityUserConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerSecurityUserConfigurationGet.json
     */
    /**
     * Sample code: Get security user configurations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSecurityUserConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserConfigurations().getWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig", com.azure.core.util.Context.NONE);
    }
}
