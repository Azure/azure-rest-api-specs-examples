
/**
 * Samples for Namespaces ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleListKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleListKey.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().listKeysWithResponse("ArunMonocle", "sdk-Namespace-2702",
            "sdk-Authrules-1746", com.azure.core.util.Context.NONE);
    }
}
