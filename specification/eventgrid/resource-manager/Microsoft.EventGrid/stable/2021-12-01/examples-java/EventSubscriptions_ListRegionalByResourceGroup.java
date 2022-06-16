import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListRegionalByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListRegionalByResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_ListRegionalByResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListRegionalByResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listRegionalByResourceGroup("examplerg", "westus2", null, null, Context.NONE);
    }
}
