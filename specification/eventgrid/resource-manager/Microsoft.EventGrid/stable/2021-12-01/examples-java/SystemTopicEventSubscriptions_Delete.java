import com.azure.core.util.Context;

/** Samples for SystemTopicEventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopicEventSubscriptions_Delete.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsDelete(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .delete("examplerg", "exampleSystemTopic1", "examplesubscription1", Context.NONE);
    }
}
