import com.azure.core.util.Context;

/** Samples for PartnerTopicEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerTopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopicEventSubscriptions()
            .getDeliveryAttributesWithResponse(
                "examplerg", "examplePartnerTopic1", "examplesubscription1", Context.NONE);
    }
}
