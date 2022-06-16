import com.azure.core.util.Context;

/** Samples for DomainTopicEventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopicEventSubscriptions_List.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainTopicEventSubscriptions()
            .list("examplerg", "exampleDomain1", "exampleDomainTopic1", null, null, Context.NONE);
    }
}
