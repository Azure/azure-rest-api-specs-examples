
/**
 * Samples for OpenShiftManagedClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2019-04-30/examples/
     * OpenShiftManagedClustersDelete.json
     */
    /**
     * Sample code: Delete OpenShift Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteOpenShiftManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getOpenShiftManagedClusters().delete("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
