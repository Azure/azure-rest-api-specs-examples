
/**
 * Samples for Namespaces GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().getAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-6914",
            "sdk-AuthRules-1788", com.azure.core.util.Context.NONE);
    }
}
