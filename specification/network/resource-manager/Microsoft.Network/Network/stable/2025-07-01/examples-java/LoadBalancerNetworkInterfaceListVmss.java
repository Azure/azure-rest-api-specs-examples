
/**
 * Samples for LoadBalancerNetworkInterfaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerNetworkInterfaceListVmss.json
     */
    /**
     * Sample code: LoadBalancerNetworkInterfaceListVmss.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerNetworkInterfaceListVmss(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerNetworkInterfaces().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
