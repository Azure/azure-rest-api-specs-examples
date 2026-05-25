
/**
 * Samples for Snapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/SnapshotsGet.json
     */
    /**
     * Sample code: Get Snapshot.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getSnapshot(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getSnapshots().getByResourceGroupWithResponse("rg1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
