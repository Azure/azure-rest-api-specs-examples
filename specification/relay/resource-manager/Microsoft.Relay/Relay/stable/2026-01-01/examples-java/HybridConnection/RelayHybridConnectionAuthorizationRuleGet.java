
/**
 * Samples for HybridConnections GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionAuthorizationRuleGet.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionAuthorizationRuleGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().getAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
