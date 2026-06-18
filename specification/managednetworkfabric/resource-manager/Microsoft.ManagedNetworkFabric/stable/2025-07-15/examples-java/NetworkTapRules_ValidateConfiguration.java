
/**
 * Samples for NetworkTapRules ValidateConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTapRules_ValidateConfiguration.json
     */
    /**
     * Sample code: NetworkTapRules_ValidateConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapRulesValidateConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTapRules().validateConfiguration("example-rg", "example-tapRule",
            com.azure.core.util.Context.NONE);
    }
}
