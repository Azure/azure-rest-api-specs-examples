
/**
 * Samples for ApplicationGateways Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayStart.json
     */
    /**
     * Sample code: Start Application Gateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void startApplicationGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().start("rg1", "appgw", com.azure.core.util.Context.NONE);
    }
}
