import com.azure.core.util.Context;

/** Samples for HybridConnections DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAutorizationRuleDelete.json
     */
    /**
     * Sample code: RelayHybridConnectionAutorizationRuleDelete.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionAutorizationRuleDelete(
        com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .deleteAuthorizationRuleWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-Hybrid-01",
                "example-RelayAuthRules-01",
                Context.NONE);
    }
}
