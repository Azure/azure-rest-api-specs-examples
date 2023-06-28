/** Samples for NetworkFabricControllers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabricControllers_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: NetworkFabricControllers_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricControllers().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
