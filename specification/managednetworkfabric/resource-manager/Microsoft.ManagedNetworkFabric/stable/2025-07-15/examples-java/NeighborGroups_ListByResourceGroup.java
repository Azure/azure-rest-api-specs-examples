
/**
 * Samples for NeighborGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NeighborGroups_ListByResourceGroup.json
     */
    /**
     * Sample code: NeighborGroups_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void neighborGroupsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.neighborGroups().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
