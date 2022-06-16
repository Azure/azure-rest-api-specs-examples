import com.azure.core.util.Context;

/** Samples for SystemTopicEventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .getWithResponse("examplerg", "exampleSystemTopic1", "examplesubscription1", Context.NONE);
    }
}
