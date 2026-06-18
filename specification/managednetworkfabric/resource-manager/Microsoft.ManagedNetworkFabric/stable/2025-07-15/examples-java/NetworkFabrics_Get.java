
/**
 * Samples for NetworkFabrics GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabrics_Get.json
     */
    /**
     * Sample code: NetworkFabrics_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabrics().getByResourceGroupWithResponse("example-rg", "example-fabric",
            com.azure.core.util.Context.NONE);
    }
}
