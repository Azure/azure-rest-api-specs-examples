
/**
 * Samples for HybridConnections ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionAuthorizationRuleListAll.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleListAll.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void
        relayHybridConnectionAuthorizationRuleListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().listAuthorizationRules("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", com.azure.core.util.Context.NONE);
    }
}
