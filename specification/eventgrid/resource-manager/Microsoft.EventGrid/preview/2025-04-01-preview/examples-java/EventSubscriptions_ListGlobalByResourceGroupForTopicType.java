
/**
 * Samples for EventSubscriptions ListGlobalByResourceGroupForTopicType.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_ListGlobalByResourceGroupForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalByResourceGroupForTopicType.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalByResourceGroupForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listGlobalByResourceGroupForTopicType("examplerg",
            "Microsoft.Resources.ResourceGroups", null, null, com.azure.core.util.Context.NONE);
    }
}
