
/**
 * Samples for ApplicationGateways ListAvailableRequestHeaders.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableRequestHeadersGet.json
     */
    /**
     * Sample code: Get Available Request Headers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableRequestHeaders(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableRequestHeadersWithResponse(com.azure.core.util.Context.NONE);
    }
}
