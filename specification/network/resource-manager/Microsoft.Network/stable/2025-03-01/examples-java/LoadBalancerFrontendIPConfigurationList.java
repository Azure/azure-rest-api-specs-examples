
/**
 * Samples for LoadBalancerFrontendIpConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LoadBalancerFrontendIPConfigurationList.json
     */
    /**
     * Sample code: LoadBalancerFrontendIPConfigurationList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerFrontendIPConfigurationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerFrontendIpConfigurations().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
