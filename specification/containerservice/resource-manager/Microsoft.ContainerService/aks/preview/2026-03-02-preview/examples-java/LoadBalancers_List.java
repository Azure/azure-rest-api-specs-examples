
/**
 * Samples for LoadBalancers ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/LoadBalancers_List.json
     */
    /**
     * Sample code: List Load Balancers by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listLoadBalancersByManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getLoadBalancers().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
