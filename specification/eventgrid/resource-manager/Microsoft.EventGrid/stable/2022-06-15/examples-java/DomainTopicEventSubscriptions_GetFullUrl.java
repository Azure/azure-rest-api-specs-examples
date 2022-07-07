import com.azure.core.util.Context;

/** Samples for DomainTopicEventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_GetFullUrl.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsGetFullUrl(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainTopicEventSubscriptions()
            .getFullUrlWithResponse(
                "examplerg", "exampleDomain1", "exampleDomainTopic1", "examplesubscription1", Context.NONE);
    }
}
