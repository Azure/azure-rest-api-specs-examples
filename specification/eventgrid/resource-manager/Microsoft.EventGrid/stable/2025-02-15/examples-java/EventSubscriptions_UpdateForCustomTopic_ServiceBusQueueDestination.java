
import com.azure.resourcemanager.eventgrid.models.EventSubscription;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.ServiceBusQueueEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.StorageBlobDeadLetterDestination;

/**
 * Samples for EventSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * EventSubscriptions_UpdateForCustomTopic_ServiceBusQueueDestination.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForCustomTopic_ServiceBusQueueDestination.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForCustomTopicServiceBusQueueDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        EventSubscription resource = manager.eventSubscriptions().getWithResponse(
            "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
            "examplesubscription1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDestination(new ServiceBusQueueEventSubscriptionDestination().withResourceId(
            "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.ServiceBus/namespaces/ContosoNamespace/queues/SBQ"))
            .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("ExamplePrefix")
                .withSubjectEndsWith("ExampleSuffix").withIsSubjectCaseSensitive(false))
            .withDeadLetterDestination(new StorageBlobDeadLetterDestination().withResourceId(
                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg")
                .withBlobContainerName("contosocontainer"))
            .apply();
    }
}
