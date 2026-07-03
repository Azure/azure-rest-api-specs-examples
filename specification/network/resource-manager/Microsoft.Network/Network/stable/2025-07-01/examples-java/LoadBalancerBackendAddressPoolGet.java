
/**
 * Samples for LoadBalancerBackendAddressPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerBackendAddressPoolGet.json
     */
    /**
     * Sample code: LoadBalancerBackendAddressPoolGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerBackendAddressPoolGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerBackendAddressPools().getWithResponse("testrg", "lb", "backend",
            com.azure.core.util.Context.NONE);
    }
}
