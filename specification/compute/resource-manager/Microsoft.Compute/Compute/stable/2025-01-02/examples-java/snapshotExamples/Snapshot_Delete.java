
/**
 * Samples for Snapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/snapshotExamples/Snapshot_Delete.json
     */
    /**
     * Sample code: Delete a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().delete("myResourceGroup", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
