/** Samples for TopicEventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: TopicEventSubscriptions_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicEventSubscriptions()
            .getWithResponse("examplerg", "exampleTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
