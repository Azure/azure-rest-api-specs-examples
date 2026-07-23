
/**
 * Samples for ManagedClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedClustersDelete.json
     */
    /**
     * Sample code: Delete Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().delete("rg1", "clustername1", null,
            com.azure.core.util.Context.NONE);
    }
}
