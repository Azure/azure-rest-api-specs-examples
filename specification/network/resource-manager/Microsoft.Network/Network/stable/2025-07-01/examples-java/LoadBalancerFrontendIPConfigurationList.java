
/**
 * Samples for LoadBalancerFrontendIpConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerFrontendIPConfigurationList.json
     */
    /**
     * Sample code: LoadBalancerFrontendIPConfigurationList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        loadBalancerFrontendIPConfigurationList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerFrontendIpConfigurations().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
