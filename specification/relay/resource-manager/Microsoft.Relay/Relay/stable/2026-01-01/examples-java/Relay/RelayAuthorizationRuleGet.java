
/**
 * Samples for WcfRelays GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayAuthorizationRuleGet.json
     */
    /**
     * Sample code: RelayAuthorizationRuleGet.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().getAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-wcf-01", "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
