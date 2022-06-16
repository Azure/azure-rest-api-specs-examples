import com.azure.core.util.Context;

/** Samples for EventSubscriptions GetFullUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_GetFullUrlForCustomTopic.json
     */
    /**
     * Sample code: EventSubscriptions_GetFullUrlForCustomTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetFullUrlForCustomTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .getFullUrlWithResponse(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
                "examplesubscription1",
                Context.NONE);
    }
}
