
/**
 * Samples for DomainTopicEventSubscriptions GetDeliveryAttributes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * DomainTopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_GetDeliveryAttributes.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopicEventSubscriptions().getDeliveryAttributesWithResponse("examplerg", "exampleDomain1",
            "exampleDomainTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
