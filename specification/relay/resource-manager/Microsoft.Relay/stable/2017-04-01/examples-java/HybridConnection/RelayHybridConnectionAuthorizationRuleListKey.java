
/**
 * Samples for HybridConnections ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/
     * RelayHybridConnectionAuthorizationRuleListKey.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleListKey.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void
        relayHybridConnectionAuthorizationRuleListKey(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().listKeysWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
