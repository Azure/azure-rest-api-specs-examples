import com.azure.core.util.Context;

/** Samples for WcfRelays ListAuthorizationRules. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleListAll.json
     */
    /**
     * Sample code: RelayAutorizationRuleListAll.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayAutorizationRuleListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .listAuthorizationRules("resourcegroup", "example-RelayNamespace-01", "example-Relay-Wcf-01", Context.NONE);
    }
}
