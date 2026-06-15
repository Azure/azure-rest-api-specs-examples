
/**
 * Samples for ManagedClusters ListClusterAdminCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ManagedClustersListClusterAdminCredentials.json
     */
    /**
     * Sample code: Get Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listClusterAdminCredentialsWithResponse("rg1", "clustername1",
            null, com.azure.core.util.Context.NONE);
    }
}
