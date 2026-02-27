
import com.azure.resourcemanager.network.fluent.models.SecurityAdminConfigurationInner;
import com.azure.resourcemanager.network.models.AddressSpaceAggregationOption;

/**
 * Samples for SecurityAdminConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * NetworkManagerSecurityAdminConfigurationPut_ManualAggregation.json
     */
    /**
     * Sample code: Create manual-mode security admin configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createManualModeSecurityAdminConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityAdminConfigurations().createOrUpdateWithResponse("rg1",
            "testNetworkManager", "myTestSecurityConfig",
            new SecurityAdminConfigurationInner()
                .withDescription("A configuration which will update any network groups ip addresses at commit times.")
                .withNetworkGroupAddressSpaceAggregationOption(AddressSpaceAggregationOption.MANUAL),
            com.azure.core.util.Context.NONE);
    }
}
