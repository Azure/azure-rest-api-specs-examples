import com.azure.core.util.Context;

/** Samples for EventHubs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubGet.json
     */
    /**
     * Sample code: EventHubGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .getWithResponse(
                "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-716", "sdk-EventHub-10", Context.NONE);
    }
}
