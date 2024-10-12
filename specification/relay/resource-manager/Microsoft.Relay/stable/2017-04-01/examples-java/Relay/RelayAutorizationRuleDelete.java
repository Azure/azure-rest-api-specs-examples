
/**
 * Samples for WcfRelays DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleDelete
     * .json
     */
    /**
     * Sample code: RelayAutorizationRuleDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAutorizationRuleDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().deleteAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-wcf-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
