
/**
 * Samples for WcfRelays DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayAuthorizationRuleDelete.json
     */
    /**
     * Sample code: RelayAuthorizationRuleDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().deleteAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-wcf-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
