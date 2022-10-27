import com.azure.core.util.Context;

/** Samples for ManagedClusters ListClusterAdminCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-09-01/examples/ManagedClustersListClusterAdminCredentials.json
     */
    /**
     * Sample code: Get Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .listClusterAdminCredentialsWithResponse("rg1", "clustername1", null, Context.NONE);
    }
}
