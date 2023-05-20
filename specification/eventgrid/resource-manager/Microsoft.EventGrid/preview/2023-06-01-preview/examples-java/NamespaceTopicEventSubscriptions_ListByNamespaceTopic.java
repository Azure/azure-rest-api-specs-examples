/** Samples for NamespaceTopicEventSubscriptions ListByNamespaceTopic. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopicEventSubscriptions_ListByNamespaceTopic.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_ListByNamespaceTopic.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicEventSubscriptionsListByNamespaceTopic(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .namespaceTopicEventSubscriptions()
            .listByNamespaceTopic(
                "examplerg",
                "examplenamespace2",
                "examplenamespacetopic2",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
