
/**
 * Samples for LoadBalancers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerList.json
     */
    /**
     * Sample code: List load balancers in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listLoadBalancersInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
