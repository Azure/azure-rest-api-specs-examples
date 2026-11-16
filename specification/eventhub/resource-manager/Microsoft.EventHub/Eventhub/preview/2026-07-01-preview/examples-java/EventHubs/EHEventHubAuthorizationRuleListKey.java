
/**
 * Samples for EventHubs ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleListKey.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleListKey.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubAuthorizationRuleListKey(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().listKeysWithResponse("ArunMonocle", "sdk-namespace-960",
            "sdk-EventHub-532", "sdk-Authrules-2513", com.azure.core.util.Context.NONE);
    }
}
