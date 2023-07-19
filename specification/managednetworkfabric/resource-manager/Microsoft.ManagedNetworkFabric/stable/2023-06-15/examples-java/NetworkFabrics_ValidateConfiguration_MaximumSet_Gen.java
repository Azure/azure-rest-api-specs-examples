import com.azure.resourcemanager.managednetworkfabric.models.ValidateAction;
import com.azure.resourcemanager.managednetworkfabric.models.ValidateConfigurationProperties;

/** Samples for NetworkFabrics ValidateConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabrics_ValidateConfiguration_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_ValidateConfiguration_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsValidateConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabrics()
            .validateConfiguration(
                "example-rg",
                "example-fabric",
                new ValidateConfigurationProperties().withValidateAction(ValidateAction.CABLING),
                com.azure.core.util.Context.NONE);
    }
}
