
/**
 * Samples for EventSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_DeleteForCustomTopic.json
     */
    /**
     * Sample code: EventSubscriptions_DeleteForCustomTopic.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        eventSubscriptionsDeleteForCustomTopic(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().delete(
            "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
