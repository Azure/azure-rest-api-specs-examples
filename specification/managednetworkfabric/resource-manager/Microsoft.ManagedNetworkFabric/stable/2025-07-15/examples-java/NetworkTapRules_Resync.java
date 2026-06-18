
/**
 * Samples for NetworkTapRules Resync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTapRules_Resync.json
     */
    /**
     * Sample code: NetworkTapRules_Resync_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapRulesResyncMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTapRules().resync("example-rg", "example-tapRule", com.azure.core.util.Context.NONE);
    }
}
