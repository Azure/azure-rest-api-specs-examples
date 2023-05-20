/** Samples for EventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_GetForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_GetForSubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetForSubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getWithResponse(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40",
                "examplesubscription3",
                com.azure.core.util.Context.NONE);
    }
}
