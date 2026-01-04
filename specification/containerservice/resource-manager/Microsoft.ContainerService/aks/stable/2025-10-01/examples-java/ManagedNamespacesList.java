
/**
 * Samples for ManagedNamespaces ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedNamespacesList.json
     */
    /**
     * Sample code: List namespaces by Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNamespacesByManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedNamespaces().listByManagedCluster("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}
