
/**
 * Samples for Namespaces DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceAuthorizationRuleDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().deleteAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-8980",
            "sdk-Authrules-8929", com.azure.core.util.Context.NONE);
    }
}
