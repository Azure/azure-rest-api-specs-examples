
/**
 * Samples for NeighborGroups Resync.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/NeighborGroups_Resync.json
     */
    /**
     * Sample code: NeighborGroups_Resync.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void
        neighborGroupsResync(com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.neighborGroups().resync("example-rg", "example-neighborgroup", com.azure.core.util.Context.NONE);
    }
}
