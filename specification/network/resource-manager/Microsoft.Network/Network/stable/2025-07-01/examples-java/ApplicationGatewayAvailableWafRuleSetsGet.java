
/**
 * Samples for ApplicationGateways ListAvailableWafRuleSets.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ApplicationGatewayAvailableWafRuleSetsGet.json
     */
    /**
     * Sample code: Get Available Waf Rule Sets.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getAvailableWafRuleSets(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getApplicationGateways()
            .listAvailableWafRuleSetsWithResponse(com.azure.core.util.Context.NONE);
    }
}
