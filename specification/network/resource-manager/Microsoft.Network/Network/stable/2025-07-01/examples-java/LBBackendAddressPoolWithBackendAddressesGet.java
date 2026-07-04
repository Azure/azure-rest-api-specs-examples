
/**
 * Samples for LoadBalancerBackendAddressPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LBBackendAddressPoolWithBackendAddressesGet.json
     */
    /**
     * Sample code: LoadBalancer with BackendAddressPool with BackendAddresses.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerWithBackendAddressPoolWithBackendAddresses(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerBackendAddressPools().getWithResponse("testrg", "lb", "backend",
            com.azure.core.util.Context.NONE);
    }
}
