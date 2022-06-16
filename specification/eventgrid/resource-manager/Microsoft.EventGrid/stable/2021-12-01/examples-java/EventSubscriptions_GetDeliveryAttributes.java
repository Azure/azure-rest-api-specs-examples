import com.azure.core.util.Context;

/** Samples for EventSubscriptions GetDeliveryAttributes. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: EventSubscriptions_GetDeliveryAttributes.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getDeliveryAttributesWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
