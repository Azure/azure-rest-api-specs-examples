
/**
 * Samples for NamespaceTopicEventSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopicEventSubscriptions_Get.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        namespaceTopicEventSubscriptionsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopicEventSubscriptions().getWithResponse("examplerg", "examplenamespace2",
            "examplenamespacetopic2", "examplenamespacetopicEventSub1", com.azure.core.util.Context.NONE);
    }
}
