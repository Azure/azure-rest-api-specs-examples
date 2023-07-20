/** Samples for NetworkTapRules ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTapRules_ListByResourceGroup_MaximumSet_Gen.json
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
