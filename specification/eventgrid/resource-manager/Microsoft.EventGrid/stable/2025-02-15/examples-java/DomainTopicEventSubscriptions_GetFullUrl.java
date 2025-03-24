
/**
 * Samples for DomainTopicEventSubscriptions GetFullUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * DomainTopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_GetFullUrl.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        domainTopicEventSubscriptionsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopicEventSubscriptions().getFullUrlWithResponse("examplerg", "exampleDomain1",
            "exampleDomainTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
