
import com.azure.resourcemanager.network.fluent.models.SecurityAdminConfigurationInner;
import com.azure.resourcemanager.network.models.AddressSpaceAggregationOption;

/**
 * Samples for SecurityAdminConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityAdminConfigurationPut_ManualAggregation.json
     */
    /**
     * Sample code: Create manual-mode security admin configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createManualModeSecurityAdminConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityAdminConfigurations().createOrUpdateWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig",
            new SecurityAdminConfigurationInner()
                .withDescription("A configuration which will update any network groups ip addresses at commit times.")
                .withNetworkGroupAddressSpaceAggregationOption(AddressSpaceAggregationOption.MANUAL),
            com.azure.core.util.Context.NONE);
    }
}
