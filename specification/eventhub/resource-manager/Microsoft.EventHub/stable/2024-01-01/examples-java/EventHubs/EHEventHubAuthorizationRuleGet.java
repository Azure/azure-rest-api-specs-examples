
/**
 * Samples for EventHubs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/
     * EHEventHubAuthorizationRuleGet.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().getAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", com.azure.core.util.Context.NONE);
    }
}
