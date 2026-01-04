
/**
 * Samples for ManagedClusters GetAccessProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersGetAccessProfile.json
     */
    /**
     * Sample code: Get Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().getAccessProfileWithResponse("rg1",
            "clustername1", "clusterUser", com.azure.core.util.Context.NONE);
    }
}
