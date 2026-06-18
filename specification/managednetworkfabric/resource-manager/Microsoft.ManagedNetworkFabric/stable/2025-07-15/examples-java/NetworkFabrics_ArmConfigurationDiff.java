
/**
 * Samples for NetworkFabrics ArmConfigurationDiff.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ArmConfigurationDiff.json
     */
    /**
     * Sample code: NetworkFabrics_ArmConfigurationDiff_MaximumSet.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsArmConfigurationDiffMaximumSet(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().armConfigurationDiff("example-rg", "example-fabric", com.azure.core.util.Context.NONE);
    }
}
