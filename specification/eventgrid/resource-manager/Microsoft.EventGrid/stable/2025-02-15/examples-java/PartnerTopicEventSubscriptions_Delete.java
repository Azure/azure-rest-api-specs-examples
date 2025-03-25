
/**
 * Samples for PartnerTopicEventSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerTopicEventSubscriptions_Delete.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerTopicEventSubscriptionsDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopicEventSubscriptions().delete("examplerg", "examplePartnerTopic1", "examplesubscription1",
            com.azure.core.util.Context.NONE);
    }
}
