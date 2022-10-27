import com.azure.core.util.Context;

/** Samples for ManagedClusters ListClusterMonitoringUserCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-09-01/examples/ManagedClustersListClusterMonitoringUserCredentials.json
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
            .listClusterMonitoringUserCredentialsWithResponse("rg1", "clustername1", null, Context.NONE);
    }
}
