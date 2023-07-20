/** Samples for NetworkFabricControllers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricControllers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabricControllers()
            .delete("example-rg", "example-networkController", com.azure.core.util.Context.NONE);
    }
}
