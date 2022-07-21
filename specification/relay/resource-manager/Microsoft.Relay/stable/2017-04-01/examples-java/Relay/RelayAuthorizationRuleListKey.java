import com.azure.core.util.Context;

/** Samples for WcfRelays ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAuthorizationRuleListKey.json
     */
    /**
     * Sample code: RelayAuthorizationRuleListKey.json.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleListKeyJson(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .listKeysWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-wcf-01",
                "example-RelayAuthRules-01",
                Context.NONE);
    }
}
