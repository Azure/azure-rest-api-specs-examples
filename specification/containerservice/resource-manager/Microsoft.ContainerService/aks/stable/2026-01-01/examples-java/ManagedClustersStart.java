
/**
 * Samples for ManagedClusters Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ManagedClustersStart.json
     */
    /**
     * Sample code: Start Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void startManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().start("rg1", "clustername1", com.azure.core.util.Context.NONE);
    }
}
