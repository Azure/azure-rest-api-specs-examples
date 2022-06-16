import com.azure.core.util.Context;

/** Samples for DomainTopicEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainTopicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainTopicEventSubscriptions()
            .getDeliveryAttributesWithResponse(
                "examplerg", "exampleDomain1", "exampleDomainTopic1", "examplesubscription1", Context.NONE);
    }
}
