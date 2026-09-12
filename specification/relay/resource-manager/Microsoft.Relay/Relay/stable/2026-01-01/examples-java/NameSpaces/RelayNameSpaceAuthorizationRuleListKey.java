
/**
 * Samples for Namespaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceAuthorizationRuleListKey.json
     */
    /**
     * Sample code: RelayNameSpaceAuthorizationRuleListKey.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().listKeysWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-RelayAuthRules-01", com.azure.core.util.Context.NONE);
    }
}
