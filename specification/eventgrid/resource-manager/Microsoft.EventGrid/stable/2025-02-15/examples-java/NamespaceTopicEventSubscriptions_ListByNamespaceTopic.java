
/**
 * Samples for NamespaceTopicEventSubscriptions ListByNamespaceTopic.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopicEventSubscriptions_ListByNamespaceTopic.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_ListByNamespaceTopic.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicEventSubscriptionsListByNamespaceTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopicEventSubscriptions().listByNamespaceTopic("examplerg", "examplenamespace2",
            "examplenamespacetopic2", null, null, com.azure.core.util.Context.NONE);
    }
}
