
/**
 * Samples for SecurityUserConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * NetworkManagerSecurityUserConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager security user configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteNetworkManagerSecurityUserConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserConfigurations().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", false, com.azure.core.util.Context.NONE);
    }
}
