
/**
 * Samples for ApplicationGateways GetSslPredefinedPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableSslOptionsPredefinedPolicyGet.json
     */
    /**
     * Sample code: Get Available Ssl Predefined Policy by name.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableSslPredefinedPolicyByName(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways().getSslPredefinedPolicyWithResponse("AppGwSslPolicy20150501",
            com.azure.core.util.Context.NONE);
    }
}
