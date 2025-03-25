
/**
 * Samples for DomainTopicEventSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
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
