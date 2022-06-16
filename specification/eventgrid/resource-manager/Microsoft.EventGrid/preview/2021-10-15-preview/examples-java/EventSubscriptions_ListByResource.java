import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_ListByResource.json
     */
    /**
     * Sample code: EventSubscriptions_ListByResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListByResource(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listByResource("examplerg", "Microsoft.EventGrid", "topics", "exampletopic2", null, null, Context.NONE);
    }
}
