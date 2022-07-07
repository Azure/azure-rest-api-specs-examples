import com.azure.core.util.Context;

/** Samples for TopicEventSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/TopicEventSubscriptions_List.json
     */
    /**
     * Sample code: TopicEventSubscriptions_List.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicEventSubscriptions().list("examplerg", "exampleTopic1", null, null, Context.NONE);
    }
}
