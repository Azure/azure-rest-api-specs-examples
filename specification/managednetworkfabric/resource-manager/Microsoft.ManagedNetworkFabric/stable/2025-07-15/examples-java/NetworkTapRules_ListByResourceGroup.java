
/**
 * Samples for NetworkTapRules ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTapRules_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkTapRules_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapRulesListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTapRules().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
