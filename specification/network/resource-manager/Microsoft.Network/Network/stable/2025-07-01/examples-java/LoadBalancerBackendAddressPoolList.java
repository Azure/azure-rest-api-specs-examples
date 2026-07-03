
/**
 * Samples for LoadBalancerBackendAddressPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerBackendAddressPoolList.json
     */
    /**
     * Sample code: LoadBalancerBackendAddressPoolList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerBackendAddressPoolList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerBackendAddressPools().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
