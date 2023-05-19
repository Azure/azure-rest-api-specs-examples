/** Samples for EventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_GetForCustomTopic_AzureFunctionDestination.json
     */
    /**
     * Sample code: EventSubscriptions_GetForCustomTopic_AzureFunctionDestination.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetForCustomTopicAzureFunctionDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getWithResponse(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
                "examplesubscription1",
                com.azure.core.util.Context.NONE);
    }
}
