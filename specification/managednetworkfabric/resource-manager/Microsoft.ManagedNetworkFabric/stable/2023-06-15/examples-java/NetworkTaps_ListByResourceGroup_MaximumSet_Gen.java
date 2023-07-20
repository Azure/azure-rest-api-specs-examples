/** Samples for NetworkTaps ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_ListByResourceGroup_MaximumSet_Gen.json
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
