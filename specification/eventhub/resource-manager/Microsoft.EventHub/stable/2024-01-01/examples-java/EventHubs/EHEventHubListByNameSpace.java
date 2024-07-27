
/**
 * Samples for EventHubs ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/
     * EHEventHubListByNameSpace.json
     */
    /**
     * Sample code: EventHubsListAll.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubsListAll(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().listByNamespace(
            "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", null, null,
            com.azure.core.util.Context.NONE);
    }
}
