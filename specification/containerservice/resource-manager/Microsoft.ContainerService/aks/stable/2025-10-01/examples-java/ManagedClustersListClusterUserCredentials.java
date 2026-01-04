
/**
 * Samples for ManagedClusters ListClusterUserCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersListClusterUserCredentials.json
     */
    /**
     * Sample code: Get Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters()
            .listClusterUserCredentialsWithResponse("rg1", "clustername1", null, null,
                com.azure.core.util.Context.NONE);
    }
}
