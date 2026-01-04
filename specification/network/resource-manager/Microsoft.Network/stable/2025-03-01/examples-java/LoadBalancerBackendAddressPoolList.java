
/**
 * Samples for LoadBalancerBackendAddressPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LoadBalancerBackendAddressPoolList.json
     */
    /**
     * Sample code: LoadBalancerBackendAddressPoolList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerBackendAddressPoolList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerBackendAddressPools().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
