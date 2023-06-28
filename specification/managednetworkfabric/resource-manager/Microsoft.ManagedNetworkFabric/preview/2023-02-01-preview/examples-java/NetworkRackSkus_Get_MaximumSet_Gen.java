/** Samples for NetworkRackSkus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkRackSkus_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkRackSkus_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkRackSkusGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkRackSkus().getWithResponse("networkRackSkuName", com.azure.core.util.Context.NONE);
    }
}
