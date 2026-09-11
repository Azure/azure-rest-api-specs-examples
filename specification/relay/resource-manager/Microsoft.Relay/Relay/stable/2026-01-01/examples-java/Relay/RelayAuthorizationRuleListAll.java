
/**
 * Samples for WcfRelays ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayAuthorizationRuleListAll.json
     */
    /**
     * Sample code: RelayAuthorizationRuleListAll.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().listAuthorizationRules("resourcegroup", "example-RelayNamespace-01", "example-Relay-Wcf-01",
            com.azure.core.util.Context.NONE);
    }
}
