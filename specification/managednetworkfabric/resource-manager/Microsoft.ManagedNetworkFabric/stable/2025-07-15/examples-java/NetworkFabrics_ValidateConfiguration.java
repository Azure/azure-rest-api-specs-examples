
import com.azure.resourcemanager.managednetworkfabric.models.ValidateAction;
import com.azure.resourcemanager.managednetworkfabric.models.ValidateConfigurationProperties;

/**
 * Samples for NetworkFabrics ValidateConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ValidateConfiguration.json
     */
    /**
     * Sample code: NetworkFabrics_ValidateConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsValidateConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().validateConfiguration("example-rg", "example-fabric",
            new ValidateConfigurationProperties().withValidateAction(ValidateAction.CABLING),
            com.azure.core.util.Context.NONE);
    }
}
