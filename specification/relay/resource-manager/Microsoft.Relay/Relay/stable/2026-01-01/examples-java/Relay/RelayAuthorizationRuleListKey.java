
/**
 * Samples for WcfRelays ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayAuthorizationRuleListKey.json
     */
    /**
     * Sample code: RelayAuthorizationRuleListKey.json.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleListKeyJson(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().listKeysWithResponse("resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01",
            "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
