
/**
 * Samples for LoadBalancers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/LoadBalancerListAll.json
     */
    /**
     * Sample code: List all load balancers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllLoadBalancers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().list(com.azure.core.util.Context.NONE);
    }
}
