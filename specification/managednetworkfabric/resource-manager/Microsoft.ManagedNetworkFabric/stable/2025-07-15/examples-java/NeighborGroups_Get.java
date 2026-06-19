
/**
 * Samples for NeighborGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NeighborGroups_Get.json
     */
    /**
     * Sample code: NeighborGroups_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void neighborGroupsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.neighborGroups().getByResourceGroupWithResponse("example-rg", "example-neighborGroup",
            com.azure.core.util.Context.NONE);
    }
}
