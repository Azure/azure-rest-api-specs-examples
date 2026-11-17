
/**
 * Samples for Namespaces GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleGet.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceAuthorizationRuleGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().getAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-2702",
            "sdk-Authrules-1746", com.azure.core.util.Context.NONE);
    }
}
