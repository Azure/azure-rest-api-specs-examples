/** Samples for NetworkFabricSkus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricSkus_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricSkus_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricSkusGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricSkus().getWithResponse("example-fabricsku", com.azure.core.util.Context.NONE);
    }
}
