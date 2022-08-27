import com.azure.core.util.Context;

/** Samples for SecurityAdminConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerSecurityAdminConfigurationGet.json
     */
    /**
     * Sample code: Get security admin configurations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSecurityAdminConfigurations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityAdminConfigurations()
            .getWithResponse("rg1", "testNetworkManager", "myTestSecurityConfig", Context.NONE);
    }
}
