/** Samples for EventSubscriptions ListByDomainTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_ListByDomainTopic.json
     */
    /**
     * Sample code: EventSubscriptions_ListByDomainTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsListByDomainTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .listByDomainTopic("examplerg", "domain1", "topic1", null, null, com.azure.core.util.Context.NONE);
    }
}
