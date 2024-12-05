
/**
 * Samples for NamespaceTopicEventSubscriptions GetFullUrl.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * NamespaceTopicEventSubscriptions_GetFullUrl.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_GetFullUrl.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        namespaceTopicEventSubscriptionsGetFullUrl(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopicEventSubscriptions().getFullUrlWithResponse("examplerg", "exampleNamespaceName1",
            "exampleDomainTopic1", "examplesubscription1", com.azure.core.util.Context.NONE);
    }
}
