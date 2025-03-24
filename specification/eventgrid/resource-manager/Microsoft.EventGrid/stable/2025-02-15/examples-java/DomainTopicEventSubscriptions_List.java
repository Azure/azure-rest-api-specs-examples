
/**
 * Samples for DomainTopicEventSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * DomainTopicEventSubscriptions_List.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_List.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopicEventSubscriptions().list("examplerg", "exampleDomain1", "exampleDomainTopic1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
