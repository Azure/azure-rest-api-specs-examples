
/**
 * Samples for Snapshots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_ListByResourceGroup.json
     */
    /**
     * Sample code: List all snapshots in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllSnapshotsInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
