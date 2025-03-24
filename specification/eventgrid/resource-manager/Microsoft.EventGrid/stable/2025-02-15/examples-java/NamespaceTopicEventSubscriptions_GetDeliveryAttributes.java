
/**
 * Samples for NamespaceTopicEventSubscriptions GetDeliveryAttributes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopicEventSubscriptions_GetDeliveryAttributes.json
     */
    /**
     * Sample code: NamespaceTopicEventSubscriptions_GetDeliveryAttributes.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicEventSubscriptionsGetDeliveryAttributes(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopicEventSubscriptions().getDeliveryAttributesWithResponse("examplerg", "exampleNamespace",
            "exampleNamespaceTopic", "exampleEventSubscriptionName", com.azure.core.util.Context.NONE);
    }
}
