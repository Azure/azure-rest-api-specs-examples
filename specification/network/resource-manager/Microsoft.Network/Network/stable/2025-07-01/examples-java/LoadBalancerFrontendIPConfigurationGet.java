
/**
 * Samples for LoadBalancerFrontendIpConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerFrontendIPConfigurationGet.json
     */
    /**
     * Sample code: LoadBalancerFrontendIPConfigurationGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        loadBalancerFrontendIPConfigurationGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerFrontendIpConfigurations().getWithResponse("testrg", "lb", "frontend",
            com.azure.core.util.Context.NONE);
    }
}
