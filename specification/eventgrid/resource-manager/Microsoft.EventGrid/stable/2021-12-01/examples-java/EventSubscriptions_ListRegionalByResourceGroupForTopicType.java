import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListRegionalByResourceGroupForTopicType. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListRegionalByResourceGroupForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListRegionalByResourceGroupForTopicType.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListRegionalByResourceGroupForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listRegionalByResourceGroupForTopicType(
                "examplerg", "westus2", "Microsoft.EventHub.namespaces", null, null, Context.NONE);
    }
}
