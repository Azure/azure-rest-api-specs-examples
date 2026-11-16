
/**
 * Samples for Namespaces ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleListAll.json
     */
    /**
     * Sample code: ListAuthorizationRules.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void listAuthorizationRules(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().listAuthorizationRules("ArunMonocle", "sdk-Namespace-2702",
            com.azure.core.util.Context.NONE);
    }
}
