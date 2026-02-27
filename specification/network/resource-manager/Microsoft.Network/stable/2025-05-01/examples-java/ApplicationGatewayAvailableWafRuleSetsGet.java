
/**
 * Samples for ApplicationGateways ListAvailableWafRuleSets.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ApplicationGatewayAvailableWafRuleSetsGet.json
     */
    /**
     * Sample code: Get Available Waf Rule Sets.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableWafRuleSets(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getApplicationGateways()
            .listAvailableWafRuleSetsWithResponse(com.azure.core.util.Context.NONE);
    }
}
