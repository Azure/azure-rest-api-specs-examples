import com.azure.core.util.Context;

/** Samples for SecurityAdminConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerSecurityAdminConfigurationList.json
     */
    /**
     * Sample code: List security admin configurations in a network manager.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSecurityAdminConfigurationsInANetworkManager(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityAdminConfigurations()
            .list("rg1", "testNetworkManager", null, null, Context.NONE);
    }
}
