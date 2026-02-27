
/**
 * Samples for LoadBalancerNetworkInterfaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * LoadBalancerNetworkInterfaceListVmss.json
     */
    /**
     * Sample code: LoadBalancerNetworkInterfaceListVmss.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerNetworkInterfaceListVmss(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerNetworkInterfaces().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
