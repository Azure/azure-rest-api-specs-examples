import com.azure.core.util.Context;

/** Samples for PartnerTopicEventSubscriptions ListByPartnerTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopicEventSubscriptions_ListByPartnerTopic.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_ListByPartnerTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicEventSubscriptionsListByPartnerTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopicEventSubscriptions()
            .listByPartnerTopic("examplerg", "examplePartnerTopic1", null, null, Context.NONE);
    }
}
