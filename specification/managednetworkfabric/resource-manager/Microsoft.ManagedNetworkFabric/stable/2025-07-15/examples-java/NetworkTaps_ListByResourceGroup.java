
/**
 * Samples for NetworkTaps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTaps_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkTaps_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTaps().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
