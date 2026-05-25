
/**
 * Samples for ManagedClusterSnapshots List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ManagedClusterSnapshotsList.json
     */
    /**
     * Sample code: List Managed Cluster Snapshots.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listManagedClusterSnapshots(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusterSnapshots().list(com.azure.core.util.Context.NONE);
    }
}
