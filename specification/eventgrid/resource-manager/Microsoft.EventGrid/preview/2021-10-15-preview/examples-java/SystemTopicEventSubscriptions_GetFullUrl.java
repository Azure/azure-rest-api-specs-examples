import com.azure.core.util.Context;

/** Samples for SystemTopicEventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/SystemTopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_GetFullUrl.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsGetFullUrl(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .getFullUrlWithResponse("examplerg", "exampleSystemTopic1", "examplesubscription1", Context.NONE);
    }
}
