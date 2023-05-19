/** Samples for PartnerTopicEventSubscriptions ListByPartnerTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerTopicEventSubscriptions_ListByPartnerTopic.json
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
            .listByPartnerTopic("examplerg", "examplePartnerTopic1", null, null, com.azure.core.util.Context.NONE);
    }
}
