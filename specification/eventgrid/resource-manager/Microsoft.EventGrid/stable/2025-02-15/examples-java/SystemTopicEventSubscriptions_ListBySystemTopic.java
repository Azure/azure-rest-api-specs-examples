
/**
 * Samples for SystemTopicEventSubscriptions ListBySystemTopic.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * SystemTopicEventSubscriptions_ListBySystemTopic.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_ListBySystemTopic.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        systemTopicEventSubscriptionsListBySystemTopic(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopicEventSubscriptions().listBySystemTopic("examplerg", "exampleSystemTopic1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
