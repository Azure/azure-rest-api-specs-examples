
/**
 * Samples for ServiceGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ServiceGatewayGet.json
     */
    /**
     * Sample code: Get load balancer.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getLoadBalancer(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().getByResourceGroupWithResponse("rg1", "sg",
            com.azure.core.util.Context.NONE);
    }
}
