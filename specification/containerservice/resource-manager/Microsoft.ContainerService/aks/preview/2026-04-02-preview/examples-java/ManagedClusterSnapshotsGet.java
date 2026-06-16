
/**
 * Samples for ManagedClusterSnapshots GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClusterSnapshotsGet.json
     */
    /**
     * Sample code: Get Managed Cluster Snapshot.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getManagedClusterSnapshot(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusterSnapshots().getByResourceGroupWithResponse("rg1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
