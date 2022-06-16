import com.azure.core.util.Context;

/** Samples for DomainEventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: DomainEventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainEventSubscriptions()
            .getDeliveryAttributesWithResponse("examplerg", "exampleDomain1", "examplesubscription1", Context.NONE);
    }
}
