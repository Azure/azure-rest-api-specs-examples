
/**
 * Samples for HybridConnections DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionAuthorizationRuleDelete.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void
        relayHybridConnectionAuthorizationRuleDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().deleteAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
