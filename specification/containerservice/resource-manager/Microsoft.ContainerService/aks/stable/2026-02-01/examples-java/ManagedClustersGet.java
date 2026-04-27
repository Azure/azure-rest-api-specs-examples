
/**
 * Samples for ManagedClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedClustersGet.json
     */
    /**
     * Sample code: Get Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getByResourceGroupWithResponse("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
