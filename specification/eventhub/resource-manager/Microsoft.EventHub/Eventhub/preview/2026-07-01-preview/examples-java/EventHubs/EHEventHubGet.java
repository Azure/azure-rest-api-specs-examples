
/**
 * Samples for EventHubs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubGet.json
     */
    /**
     * Sample code: EventHubGet.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubGet(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().getWithResponse("Default-NotificationHubs-AustraliaEast",
            "sdk-Namespace-716", "sdk-EventHub-10", com.azure.core.util.Context.NONE);
    }
}
