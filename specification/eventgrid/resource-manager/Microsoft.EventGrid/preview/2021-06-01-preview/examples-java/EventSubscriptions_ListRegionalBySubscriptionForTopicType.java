import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListRegionalBySubscriptionForTopicType. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventSubscriptions_ListRegionalBySubscriptionForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListRegionalBySubscriptionForTopicType.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListRegionalBySubscriptionForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listRegionalBySubscriptionForTopicType(
                "westus2", "Microsoft.EventHub.namespaces", null, null, Context.NONE);
    }
}
