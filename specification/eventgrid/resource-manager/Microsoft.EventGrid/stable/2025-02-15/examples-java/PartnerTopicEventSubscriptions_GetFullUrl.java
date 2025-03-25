
/**
 * Samples for PartnerTopicEventSubscriptions GetFullUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerTopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: PartnerTopicEventSubscriptions_GetFullUrl.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        partnerTopicEventSubscriptionsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerTopicEventSubscriptions().getFullUrlWithResponse("examplerg", "examplePartnerTopic1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
