
/**
 * Samples for Namespaces ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceAuthorizationRuleListAll.json
     */
    /**
     * Sample code: RelayNameSpaceAuthorizationRuleListAll.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAuthorizationRuleListAll(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().listAuthorizationRules("resourcegroup", "example-RelayNamespace-01",
            com.azure.core.util.Context.NONE);
    }
}
