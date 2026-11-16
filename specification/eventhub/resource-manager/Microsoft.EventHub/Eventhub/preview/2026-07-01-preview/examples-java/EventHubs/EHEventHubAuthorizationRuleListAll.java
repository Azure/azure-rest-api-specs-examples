
/**
 * Samples for EventHubs ListAuthorizationRules.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleListAll.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleListAll.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubAuthorizationRuleListAll(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().listAuthorizationRules("ArunMonocle", "sdk-Namespace-960",
            "sdk-EventHub-532", com.azure.core.util.Context.NONE);
    }
}
