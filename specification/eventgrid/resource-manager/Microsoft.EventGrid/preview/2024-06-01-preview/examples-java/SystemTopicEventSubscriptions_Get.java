
/**
 * Samples for SystemTopicEventSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * SystemTopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.systemTopicEventSubscriptions().getWithResponse("examplerg", "exampleSystemTopic1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
