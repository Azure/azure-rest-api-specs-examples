
/**
 * Samples for DomainEventSubscriptions GetDeliveryAttributes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * DomainEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: DomainEventSubscriptions_GetDeliveryAttributes.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        domainEventSubscriptionsGetDeliveryAttributes(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().getDeliveryAttributesWithResponse("examplerg", "exampleDomain1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
