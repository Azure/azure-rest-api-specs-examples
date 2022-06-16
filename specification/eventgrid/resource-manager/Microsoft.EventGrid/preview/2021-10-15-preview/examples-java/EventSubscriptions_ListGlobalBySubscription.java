import com.azure.core.util.Context;

/** Samples for EventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_ListGlobalBySubscription.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalBySubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalBySubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().list(null, null, Context.NONE);
    }
}
