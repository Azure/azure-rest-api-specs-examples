
import com.azure.resourcemanager.network.fluent.models.SecurityUserConfigurationInner;

/**
 * Samples for SecurityUserConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserConfigurationPut.json
     */
    /**
     * Sample code: Create network manager security user configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createNetworkManagerSecurityUserConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserConfigurations().createOrUpdateWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", new SecurityUserConfigurationInner().withDescription("A sample policy"),
            com.azure.core.util.Context.NONE);
    }
}
