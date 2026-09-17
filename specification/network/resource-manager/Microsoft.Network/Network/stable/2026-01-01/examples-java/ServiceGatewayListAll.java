
/**
 * Samples for ServiceGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ServiceGatewayListAll.json
     */
    /**
     * Sample code: List all load balancers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllLoadBalancers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getServiceGateways().list(com.azure.core.util.Context.NONE);
    }
}
