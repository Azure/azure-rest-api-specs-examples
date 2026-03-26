
/**
 * Samples for ManagedNamespaces ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ManagedNamespacesList.json
     */
    /**
     * Sample code: List namespaces by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listNamespacesByManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedNamespaces().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
