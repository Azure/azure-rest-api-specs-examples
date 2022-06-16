import com.azure.core.util.Context;

/** Samples for EventSubscriptions ListByDomainTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_ListByDomainTopic.json
     */
    /**
     * Sample code: EventSubscriptions_ListByDomainTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListByDomainTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().listByDomainTopic("examplerg", "domain1", "topic1", null, null, Context.NONE);
    }
}
