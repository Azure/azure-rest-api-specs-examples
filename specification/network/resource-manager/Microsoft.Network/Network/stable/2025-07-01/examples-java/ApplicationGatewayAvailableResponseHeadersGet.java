
/**
 * Samples for ApplicationGateways ListAvailableResponseHeaders.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableResponseHeadersGet.json
     */
    /**
     * Sample code: Get Available Response Headers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableResponseHeaders(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableResponseHeadersWithResponse(com.azure.core.util.Context.NONE);
    }
}
