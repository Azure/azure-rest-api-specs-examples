
/**
 * Samples for ServiceGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayGet.json
     */
    /**
     * Sample code: Get load balancer.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLoadBalancer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getServiceGateways().getByResourceGroupWithResponse("rg1", "sg",
            com.azure.core.util.Context.NONE);
    }
}
