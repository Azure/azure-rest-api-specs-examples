import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListGlobalByResourceGroupForTopicType. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListGlobalByResourceGroupForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalByResourceGroupForTopicType.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalByResourceGroupForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listGlobalByResourceGroupForTopicType(
                "examplerg", "Microsoft.Resources.ResourceGroups", null, null, Context.NONE);
    }
}
