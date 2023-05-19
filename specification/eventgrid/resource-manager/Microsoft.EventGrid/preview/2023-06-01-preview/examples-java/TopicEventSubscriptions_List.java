/** Samples for TopicEventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicEventSubscriptions_List.json
     */
    /**
     * Sample code: TopicEventSubscriptions_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicEventSubscriptions()
            .list("examplerg", "exampleTopic1", null, null, com.azure.core.util.Context.NONE);
    }
}
