
/**
 * Samples for NetworkFabricControllers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabricControllers_Delete.json
     */
    /**
     * Sample code: NetworkFabricControllers_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricControllers().delete("example-rg", "example-networkController",
            com.azure.core.util.Context.NONE);
    }
}
