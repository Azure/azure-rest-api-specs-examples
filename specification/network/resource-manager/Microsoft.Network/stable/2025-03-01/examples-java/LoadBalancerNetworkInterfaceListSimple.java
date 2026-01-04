
/**
 * Samples for LoadBalancerNetworkInterfaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LoadBalancerNetworkInterfaceListSimple.json
     */
    /**
     * Sample code: LoadBalancerNetworkInterfaceListSimple.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerNetworkInterfaceListSimple(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerNetworkInterfaces().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
