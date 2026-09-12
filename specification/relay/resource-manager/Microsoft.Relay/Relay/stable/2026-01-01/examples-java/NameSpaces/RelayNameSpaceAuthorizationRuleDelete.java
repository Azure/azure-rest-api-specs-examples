
/**
 * Samples for Namespaces DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceAuthorizationRuleDelete.json
     */
    /**
     * Sample code: RelayNameSpaceAuthorizationRuleDelete.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAuthorizationRuleDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().deleteAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
