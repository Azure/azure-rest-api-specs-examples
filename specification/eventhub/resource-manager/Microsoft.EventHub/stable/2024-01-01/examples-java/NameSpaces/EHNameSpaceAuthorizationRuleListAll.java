
/**
 * Samples for Namespaces ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/
     * EHNameSpaceAuthorizationRuleListAll.json
     */
    /**
     * Sample code: ListAuthorizationRules.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuthorizationRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().listAuthorizationRules("ArunMonocle",
            "sdk-Namespace-2702", com.azure.core.util.Context.NONE);
    }
}
