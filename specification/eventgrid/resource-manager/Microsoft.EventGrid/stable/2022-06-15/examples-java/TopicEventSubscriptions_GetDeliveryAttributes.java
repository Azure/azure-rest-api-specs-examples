import com.azure.core.util.Context;

/** Samples for TopicEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/TopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: TopicEventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicEventSubscriptions()
            .getDeliveryAttributesWithResponse("examplerg", "exampleTopic1", "examplesubscription1", Context.NONE);
    }
}
