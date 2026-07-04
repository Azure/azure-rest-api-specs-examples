
/**
 * Samples for ApplicationGateways BackendHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayBackendHealthGet.json
     */
    /**
     * Sample code: Get Backend Health.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getBackendHealth(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().backendHealth("appgw", "appgw", null,
            com.azure.core.util.Context.NONE);
    }
}
