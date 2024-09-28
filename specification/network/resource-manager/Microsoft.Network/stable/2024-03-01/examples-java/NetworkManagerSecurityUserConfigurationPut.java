
import com.azure.resourcemanager.network.fluent.models.SecurityUserConfigurationInner;

/**
 * Samples for SecurityUserConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/
     * NetworkManagerSecurityUserConfigurationPut.json
     */
    /**
     * Sample code: Create network manager security user configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createNetworkManagerSecurityUserConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityUserConfigurations().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig",
            new SecurityUserConfigurationInner().withDescription("A sample policy"), com.azure.core.util.Context.NONE);
    }
}
