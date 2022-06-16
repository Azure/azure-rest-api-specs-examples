import com.azure.core.util.Context;

/** Samples for EventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventSubscriptions_GetFullUrlForResource.json
     */
    /**
     * Sample code: EventSubscriptions_GetFullUrlForResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetFullUrlForResource(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getFullUrlWithResponse(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1",
                "examplesubscription1",
                Context.NONE);
    }
}
