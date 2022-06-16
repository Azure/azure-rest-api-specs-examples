import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListGlobalBySubscriptionForTopicType. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListGlobalBySubscriptionForTopicType.json
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
