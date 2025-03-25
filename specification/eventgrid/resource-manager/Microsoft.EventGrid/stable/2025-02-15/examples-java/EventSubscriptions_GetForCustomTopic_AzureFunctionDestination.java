
/**
 * Samples for EventSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_GetForCustomTopic_AzureFunctionDestination.json
     */
    /**
     * Sample code: EventSubscriptions_GetForCustomTopic_AzureFunctionDestination.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsGetForCustomTopicAzureFunctionDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().getWithResponse(
            "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
            "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
