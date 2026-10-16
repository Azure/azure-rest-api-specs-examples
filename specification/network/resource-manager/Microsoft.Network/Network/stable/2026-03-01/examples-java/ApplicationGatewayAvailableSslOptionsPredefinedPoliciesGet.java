
/**
 * Samples for ApplicationGateways ListAvailableSslPredefinedPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ApplicationGatewayAvailableSslOptionsPredefinedPoliciesGet.json
     */
    /**
     * Sample code: Get Available Ssl Predefined Policies.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableSslPredefinedPolicies(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableSslPredefinedPolicies(com.azure.core.util.Context.NONE);
    }
}
