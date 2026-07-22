
/**
 * Samples for ManagedClusters ListClusterUserCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedClustersListClusterUserCredentials.json
     */
    /**
     * Sample code: Get Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listClusterUserCredentialsWithResponse("rg1", "clustername1", null,
            null, com.azure.core.util.Context.NONE);
    }
}
