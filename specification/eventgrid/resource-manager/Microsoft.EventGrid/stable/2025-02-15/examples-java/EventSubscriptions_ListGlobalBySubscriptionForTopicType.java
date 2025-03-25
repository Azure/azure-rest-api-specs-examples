
/**
 * Samples for EventSubscriptions ListGlobalBySubscriptionForTopicType.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_ListGlobalBySubscriptionForTopicType.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalBySubscriptionForTopicType.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalBySubscriptionForTopicType(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listGlobalBySubscriptionForTopicType("Microsoft.Resources.Subscriptions", null,
            null, com.azure.core.util.Context.NONE);
    }
}
