
/**
 * Samples for ManagedClusterSnapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClusterSnapshotsDelete.json
     */
    /**
     * Sample code: Delete Managed Cluster Snapshot.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteManagedClusterSnapshot(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusterSnapshots().deleteWithResponse("rg1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
