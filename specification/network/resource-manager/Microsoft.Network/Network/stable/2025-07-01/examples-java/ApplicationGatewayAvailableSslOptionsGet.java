
/**
 * Samples for ApplicationGateways ListAvailableSslOptions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableSslOptionsGet.json
     */
    /**
     * Sample code: Get Available Ssl Options.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableSslOptions(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableSslOptionsWithResponse(com.azure.core.util.Context.NONE);
    }
}
