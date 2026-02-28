
/**
 * Samples for LoadBalancerBackendAddressPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * LoadBalancerBackendAddressPoolDelete.json
     */
    /**
     * Sample code: BackendAddressPoolDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void backendAddressPoolDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerBackendAddressPools().delete("testrg", "lb",
            "backend", com.azure.core.util.Context.NONE);
    }
}
