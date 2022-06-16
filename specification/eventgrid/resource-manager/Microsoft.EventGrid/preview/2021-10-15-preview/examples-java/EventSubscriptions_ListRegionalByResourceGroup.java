import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListRegionalByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_ListRegionalByResourceGroup.json
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
