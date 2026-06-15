
/**
 * Samples for ManagedClusters Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClustersStop.json
     */
    /**
     * Sample code: Stop Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void stopManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().stop("rg1", "clustername1", com.azure.core.util.Context.NONE);
    }
}
