import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListGlobalBySubscriptionForTopicType. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventSubscriptions_ListGlobalBySubscriptionForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalBySubscriptionForTopicType.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalBySubscriptionForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listGlobalBySubscriptionForTopicType("Microsoft.Resources.Subscriptions", null, null, Context.NONE);
    }
}
