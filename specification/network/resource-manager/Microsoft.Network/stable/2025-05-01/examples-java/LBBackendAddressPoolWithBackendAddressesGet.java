
/**
 * Samples for LoadBalancerBackendAddressPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * LBBackendAddressPoolWithBackendAddressesGet.json
     */
    /**
     * Sample code: LoadBalancer with BackendAddressPool with BackendAddresses.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        loadBalancerWithBackendAddressPoolWithBackendAddresses(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerBackendAddressPools().getWithResponse("testrg", "lb",
            "backend", com.azure.core.util.Context.NONE);
    }
}
