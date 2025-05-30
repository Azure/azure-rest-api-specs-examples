
import com.azure.resourcemanager.network.fluent.models.SecurityAdminConfigurationInner;
import com.azure.resourcemanager.network.models.NetworkIntentPolicyBasedService;
import java.util.Arrays;

/**
 * Samples for SecurityAdminConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * NetworkManagerSecurityAdminConfigurationPut.json
     */
    /**
     * Sample code: Create network manager security admin configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createNetworkManagerSecurityAdminConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityAdminConfigurations().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig",
            new SecurityAdminConfigurationInner().withDescription("A sample policy")
                .withApplyOnNetworkIntentPolicyBasedServices(Arrays.asList(NetworkIntentPolicyBasedService.NONE)),
            com.azure.core.util.Context.NONE);
    }
}
