
/**
 * Samples for EventSubscriptions ListRegionalBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_ListRegionalBySubscription.json
     */
    /**
     * Sample code: EventSubscriptions_ListRegionalBySubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        eventSubscriptionsListRegionalBySubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listRegionalBySubscription("westus2", null, null,
            com.azure.core.util.Context.NONE);
    }
}
