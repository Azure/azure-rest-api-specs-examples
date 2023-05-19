/** Samples for SystemTopicEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/SystemTopicEventSubscriptions_GetDeliveryAttributes.json
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
                "examplerg", "exampleSystemTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
