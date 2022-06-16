import com.azure.core.util.Context;

/** Samples for TopicEventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/TopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: TopicEventSubscriptions_GetFullUrl.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicEventSubscriptions()
            .getFullUrlWithResponse("examplerg", "exampleTopic1", "examplesubscription1", Context.NONE);
    }
}
