
/**
 * Samples for Snapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/SnapshotsDelete.json
     */
    /**
     * Sample code: Delete Snapshot.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void deleteSnapshot(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getSnapshots().deleteWithResponse("rg1", "snapshot1", com.azure.core.util.Context.NONE);
    }
}
