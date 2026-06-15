
/**
 * Samples for ManagedClusterSnapshots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClusterSnapshotsListByResourceGroup.json
     */
    /**
     * Sample code: List Managed Cluster Snapshots by Resource Group.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listManagedClusterSnapshotsByResourceGroup(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusterSnapshots().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
