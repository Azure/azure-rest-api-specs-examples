
/**
 * Samples for LoadBalancerNetworkInterfaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerNetworkInterfaceListSimple.json
     */
    /**
     * Sample code: LoadBalancerNetworkInterfaceListSimple.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        loadBalancerNetworkInterfaceListSimple(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerNetworkInterfaces().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
