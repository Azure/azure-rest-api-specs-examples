
/**
 * Samples for EventHubs GetAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleGet.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubAuthorizationRuleGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().getAuthorizationRuleWithResponse("ArunMonocle", "sdk-Namespace-960",
            "sdk-EventHub-532", "sdk-Authrules-2513", com.azure.core.util.Context.NONE);
    }
}
