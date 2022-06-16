import com.azure.core.util.Context;

/** Samples for EventSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_DeleteForResource.json
     */
    /**
     * Sample code: EventSubscriptions_DeleteForResource.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsDeleteForResource(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .delete(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1",
                "examplesubscription10",
                Context.NONE);
    }
}
