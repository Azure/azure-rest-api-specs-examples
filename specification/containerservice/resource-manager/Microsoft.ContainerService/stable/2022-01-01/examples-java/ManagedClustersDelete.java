import com.azure.core.util.Context;

/** Samples for ManagedClusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-01-01/examples/ManagedClustersDelete.json
     */
    /**
     * Sample code: Delete Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .delete("rg1", "clustername1", Context.NONE);
    }
}
