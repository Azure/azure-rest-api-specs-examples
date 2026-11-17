
/**
 * Samples for EventHubs DeleteAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleDelete.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubAuthorizationRuleDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().deleteAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-960",
            "sdk-EventHub-532", "sdk-Authrules-2513", com.azure.core.util.Context.NONE);
    }
}
