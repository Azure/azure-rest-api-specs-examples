
/**
 * Samples for NetworkFabrics ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkFabrics_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
