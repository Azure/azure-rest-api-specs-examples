/** Samples for SystemTopicEventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/SystemTopicEventSubscriptions_Delete.json
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
            .delete("examplerg", "exampleSystemTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
