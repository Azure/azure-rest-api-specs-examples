
/**
 * Samples for Namespaces ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleListAll.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListAll.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleListAll(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().listAuthorizationRules("ArunMonocle", "sdk-Namespace-6914",
            com.azure.core.util.Context.NONE);
    }
}
