
/**
 * Samples for EventSubscriptions GetFullUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_GetFullUrlForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_GetFullUrlForSubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        eventSubscriptionsGetFullUrlForSubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().getFullUrlWithResponse("subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4",
            "examplesubscription3", com.azure.core.util.Context.NONE);
    }
}
