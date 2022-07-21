import com.azure.core.util.Context;

/** Samples for WcfRelays GetAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleGet.json
     */
    /**
     * Sample code: RelayAutorizationRuleGet.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayAutorizationRuleGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .getAuthorizationRuleWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-wcf-01",
                "example-RelayAuthRules-01",
                Context.NONE);
    }
}
