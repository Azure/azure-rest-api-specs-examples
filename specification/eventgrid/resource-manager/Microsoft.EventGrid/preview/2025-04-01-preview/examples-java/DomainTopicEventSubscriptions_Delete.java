
/**
 * Samples for DomainTopicEventSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * DomainTopicEventSubscriptions_Delete.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        domainTopicEventSubscriptionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopicEventSubscriptions().delete("examplerg", "exampleDomain1", "exampleDomainTopic1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
