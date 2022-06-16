import com.azure.core.util.Context;

/** Samples for PartnerTopicEventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerTopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_Get.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopicEventSubscriptions()
            .getWithResponse("examplerg", "examplePartnerTopic1", "examplesubscription1", Context.NONE);
    }
}
