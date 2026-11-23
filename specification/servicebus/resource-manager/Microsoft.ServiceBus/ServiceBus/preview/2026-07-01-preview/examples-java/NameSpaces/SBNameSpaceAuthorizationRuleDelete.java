
/**
 * Samples for Namespaces DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleDelete.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().deleteAuthorizationRuleWithResponse("ArunMonocle", "sdk-namespace-6914",
            "sdk-AuthRules-1788", com.azure.core.util.Context.NONE);
    }
}
