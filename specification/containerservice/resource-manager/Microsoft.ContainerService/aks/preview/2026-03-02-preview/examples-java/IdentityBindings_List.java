
/**
 * Samples for IdentityBindings ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/IdentityBindings_List.json
     */
    /**
     * Sample code: List Identity Bindings by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listIdentityBindingsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getIdentityBindings().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
