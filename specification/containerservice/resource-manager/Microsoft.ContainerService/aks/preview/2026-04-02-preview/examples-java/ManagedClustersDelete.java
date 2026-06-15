
/**
 * Samples for ManagedClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClustersDelete.json
     */
    /**
     * Sample code: Delete Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().delete("rg1", "clustername1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
