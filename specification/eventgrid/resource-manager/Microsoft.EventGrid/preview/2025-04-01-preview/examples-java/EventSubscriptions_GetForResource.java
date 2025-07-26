
/**
 * Samples for EventSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_GetForResource.json
     */
    /**
     * Sample code: EventSubscriptions_GetForResource.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetForResource(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().getWithResponse(
            "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
