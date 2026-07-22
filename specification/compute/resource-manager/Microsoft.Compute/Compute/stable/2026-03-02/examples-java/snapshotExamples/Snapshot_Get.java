
/**
 * Samples for Snapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/snapshotExamples/Snapshot_Get.json
     */
    /**
     * Sample code: Get information about a snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutASnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().getByResourceGroupWithResponse("myResourceGroup", "mySnapshot",
            com.azure.core.util.Context.NONE);
    }
}
