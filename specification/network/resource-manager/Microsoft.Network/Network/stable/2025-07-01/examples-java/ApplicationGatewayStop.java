
/**
 * Samples for ApplicationGateways Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayStop.json
     */
    /**
     * Sample code: Stop Application Gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void stopApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().stop("rg1", "appgw", com.azure.core.util.Context.NONE);
    }
}
