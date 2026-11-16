
/**
 * Samples for EventHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubDelete.json
     */
    /**
     * Sample code: EventHubDelete.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubDelete(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().deleteWithResponse("ArunMonocle", "sdk-Namespace-5357",
            "sdk-EventHub-6547", com.azure.core.util.Context.NONE);
    }
}
