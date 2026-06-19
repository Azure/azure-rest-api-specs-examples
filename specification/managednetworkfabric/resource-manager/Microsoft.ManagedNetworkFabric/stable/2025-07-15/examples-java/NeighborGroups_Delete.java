
/**
 * Samples for NeighborGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NeighborGroups_Delete.json
     */
    /**
     * Sample code: NeighborGroups_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void neighborGroupsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.neighborGroups().delete("example-rg", "example-neighborGroup", com.azure.core.util.Context.NONE);
    }
}
