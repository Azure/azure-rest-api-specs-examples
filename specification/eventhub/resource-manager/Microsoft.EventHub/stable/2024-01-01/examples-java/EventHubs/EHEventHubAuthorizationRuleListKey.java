
/**
 * Samples for EventHubs ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/
     * EHEventHubAuthorizationRuleListKey.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleListKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleListKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().listKeysWithResponse("ArunMonocle",
            "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", com.azure.core.util.Context.NONE);
    }
}
