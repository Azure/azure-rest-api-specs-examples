/** Samples for NetworkFabrics Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabrics_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .networkFabrics()
            .delete("rgNetworkFabrics", "lrhjxlxlhgvufessdcuetcwnto", com.azure.core.util.Context.NONE);
    }
}
