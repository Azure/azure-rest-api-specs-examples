import com.azure.core.util.Context;

/** Samples for PartnerTopicEventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/PartnerTopicEventSubscriptions_Delete.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_Delete.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicEventSubscriptionsDelete(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopicEventSubscriptions()
            .delete("examplerg", "examplePartnerTopic1", "examplesubscription1", Context.NONE);
    }
}
