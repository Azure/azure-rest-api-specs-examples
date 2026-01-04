
/**
 * Samples for LoadBalancerBackendAddressPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LBBackendAddressPoolListWithBackendAddressesPoolType.json
     */
    /**
     * Sample code: Load balancer with BackendAddressPool containing BackendAddresses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerWithBackendAddressPoolContainingBackendAddresses(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerBackendAddressPools().list("testrg", "lb",
            com.azure.core.util.Context.NONE);
    }
}
