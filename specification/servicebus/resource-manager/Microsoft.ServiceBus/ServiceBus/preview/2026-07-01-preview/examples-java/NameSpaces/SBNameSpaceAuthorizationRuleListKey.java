
/**
 * Samples for Namespaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleListKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().listKeysWithResponse("ArunMonocle", "sdk-namespace-6914",
            "sdk-AuthRules-1788", com.azure.core.util.Context.NONE);
    }
}
