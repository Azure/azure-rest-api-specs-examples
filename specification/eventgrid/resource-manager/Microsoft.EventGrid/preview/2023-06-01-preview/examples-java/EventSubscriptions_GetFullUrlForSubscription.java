/** Samples for EventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_GetFullUrlForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_GetFullUrlForSubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetFullUrlForSubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getFullUrlWithResponse(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40",
                "examplesubscription3",
                com.azure.core.util.Context.NONE);
    }
}
