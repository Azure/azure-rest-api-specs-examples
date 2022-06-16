import com.azure.core.util.Context;

/** Samples for EventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_GetFullUrlForResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_GetFullUrlForResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetFullUrlForResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getFullUrlWithResponse(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg",
                "examplesubscription2",
                Context.NONE);
    }
}
