import com.azure.core.util.Context;

/** Samples for SecurityAdminConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkManagerSecurityAdminConfigurationDelete.json
     */
    /**
     * Sample code: Delete network manager security admin configuration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkManagerSecurityAdminConfiguration(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityAdminConfigurations()
            .delete("rg1", "testNetworkManager", "myTestSecurityConfig", false, Context.NONE);
    }
}
