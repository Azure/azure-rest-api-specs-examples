
/**
 * Samples for LoadBalancers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/LoadBalancers_Delete.json
     */
    /**
     * Sample code: Delete a Load Balancer.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void deleteALoadBalancer(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getLoadBalancers().delete("rg1", "clustername1", "kubernetes",
            com.azure.core.util.Context.NONE);
    }
}
