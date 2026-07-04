
import com.azure.resourcemanager.network.fluent.models.SecurityAdminConfigurationInner;
import com.azure.resourcemanager.network.models.NetworkIntentPolicyBasedService;
import java.util.Arrays;

/**
 * Samples for SecurityAdminConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationPut.json
     */
    /**
     * Sample code: Create network manager security admin configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createNetworkManagerSecurityAdminConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityAdminConfigurations().createOrUpdateWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig",
            new SecurityAdminConfigurationInner().withDescription("A sample policy")
                .withApplyOnNetworkIntentPolicyBasedServices(Arrays.asList(NetworkIntentPolicyBasedService.NONE)),
            com.azure.core.util.Context.NONE);
    }
}
