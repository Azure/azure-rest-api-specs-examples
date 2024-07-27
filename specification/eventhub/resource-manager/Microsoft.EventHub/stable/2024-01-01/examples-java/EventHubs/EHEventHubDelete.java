
/**
 * Samples for EventHubs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubDelete.
     * json
     */
    /**
     * Sample code: EventHubDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().deleteWithResponse("ArunMonocle",
            "sdk-Namespace-5357", "sdk-EventHub-6547", com.azure.core.util.Context.NONE);
    }
}
