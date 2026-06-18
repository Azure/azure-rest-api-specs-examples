
/**
 * Samples for NetworkTapRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTapRules_Delete.json
     */
    /**
     * Sample code: NetworkTapRules_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapRulesDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTapRules().delete("example-rg", "example-tapRule", com.azure.core.util.Context.NONE);
    }
}
