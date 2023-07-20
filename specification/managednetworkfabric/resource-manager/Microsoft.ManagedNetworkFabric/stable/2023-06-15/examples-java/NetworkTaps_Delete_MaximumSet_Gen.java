/** Samples for NetworkTaps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkTaps_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkTapsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkTaps().delete("example-rg", "example-networkTap", com.azure.core.util.Context.NONE);
    }
}
