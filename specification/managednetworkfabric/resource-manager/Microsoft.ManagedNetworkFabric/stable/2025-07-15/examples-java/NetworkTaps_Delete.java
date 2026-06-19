
/**
 * Samples for NetworkTaps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkTaps_Delete.json
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
