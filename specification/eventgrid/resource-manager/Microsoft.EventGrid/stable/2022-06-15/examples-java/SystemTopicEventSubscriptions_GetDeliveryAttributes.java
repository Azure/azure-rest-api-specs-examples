import com.azure.core.util.Context;

/** Samples for SystemTopicEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .getDeliveryAttributesWithResponse(
                "examplerg", "exampleSystemTopic1", "examplesubscription1", Context.NONE);
    }
}
