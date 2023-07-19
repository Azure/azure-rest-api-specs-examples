/** Samples for NetworkTaps Resync. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_Resync_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkTaps_Resync_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapsResyncMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTaps().resync("example-rg", "example-networkTap", com.azure.core.util.Context.NONE);
    }
}
