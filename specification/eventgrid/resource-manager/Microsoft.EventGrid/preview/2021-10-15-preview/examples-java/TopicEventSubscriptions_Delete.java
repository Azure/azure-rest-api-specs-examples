import com.azure.core.util.Context;

/** Samples for TopicEventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/TopicEventSubscriptions_Delete.json
     */
    /**
     * Sample code: TopicEventSubscriptions_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicEventSubscriptions().delete("examplerg", "exampleTopic1", "examplesubscription1", Context.NONE);
    }
}
