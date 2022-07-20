import com.azure.core.util.Context;

/** Samples for HybridConnections ListAuthorizationRules. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAutorizationRuleListAll.json
     */
    /**
     * Sample code: RelayHybridConnectionAutorizationRuleListAll.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionAutorizationRuleListAll(
        com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .listAuthorizationRules(
                "resourcegroup", "example-RelayNamespace-01", "example-Relay-Hybrid-01", Context.NONE);
    }
}
