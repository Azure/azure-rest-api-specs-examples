
/**
 * Samples for NetworkFabricControllers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NetworkFabricControllers_ListByResourceGroup.json
     */
    /**
     * Sample code: NetworkFabricControllers_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void networkFabricControllersListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.networkFabricControllers().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
