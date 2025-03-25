
/**
 * Samples for EventSubscriptions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_ListGlobalByResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalByResourceGroup.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        eventSubscriptionsListGlobalByResourceGroup(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listByResourceGroup("examplerg", null, null, com.azure.core.util.Context.NONE);
    }
}
