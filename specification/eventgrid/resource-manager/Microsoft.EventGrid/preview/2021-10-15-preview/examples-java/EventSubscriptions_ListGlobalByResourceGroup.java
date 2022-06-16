import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_ListGlobalByResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_ListGlobalByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListGlobalByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listByResourceGroup("examplerg", null, null, Context.NONE);
    }
}
