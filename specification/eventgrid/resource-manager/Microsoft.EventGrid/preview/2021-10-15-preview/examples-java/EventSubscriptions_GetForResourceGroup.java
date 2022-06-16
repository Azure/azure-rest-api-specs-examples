import com.azure.core.util.Context;

/** Samples for EventSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_GetForResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_GetForResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetForResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getWithResponse(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg",
                "examplesubscription2",
                Context.NONE);
    }
}
