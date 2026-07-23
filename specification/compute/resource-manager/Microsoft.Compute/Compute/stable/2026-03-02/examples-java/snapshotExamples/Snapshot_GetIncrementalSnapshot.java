
/**
 * Samples for Snapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_GetIncrementalSnapshot.json
     */
    /**
     * Sample code: Get information about an incremental snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getInformationAboutAnIncrementalSnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().getByResourceGroupWithResponse("myResourceGroup",
            "myIncrementalSnapshot", com.azure.core.util.Context.NONE);
    }
}
