
/**
 * Samples for EventSubscriptions ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * EventSubscriptions_ListByResource.json
     */
    /**
     * Sample code: EventSubscriptions_ListByResource.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListByResource(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listByResource("examplerg", "Microsoft.EventGrid", "topics", "exampletopic2", null,
            null, com.azure.core.util.Context.NONE);
    }
}
