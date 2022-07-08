import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListRegionalBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_ListRegionalBySubscription.json
     */
    /**
     * Sample code: EventSubscriptions_ListRegionalBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListRegionalBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listRegionalBySubscription("westus2", null, null, Context.NONE);
    }
}
