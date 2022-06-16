import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.HybridConnectionEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.StorageBlobDeadLetterDestination;

/** Samples for EventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-06-01-preview/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_HybridConnectionDestination.json
     */
    /**
     * Sample code: EventSubscriptions_CreateOrUpdateForCustomTopic_HybridConnectionDestination.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsCreateOrUpdateForCustomTopicHybridConnectionDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .define("examplesubscription1")
            .withExistingScope(
                "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1")
            .withDestination(
                new HybridConnectionEventSubscriptionDestination()
                    .withResourceId(
                        "/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1"))
            .withFilter(
                new EventSubscriptionFilter()
                    .withSubjectBeginsWith("ExamplePrefix")
                    .withSubjectEndsWith("ExampleSuffix")
                    .withIsSubjectCaseSensitive(false))
            .withDeadLetterDestination(
                new StorageBlobDeadLetterDestination()
                    .withResourceId(
                        "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg")
                    .withBlobContainerName("contosocontainer"))
            .create();
    }
}
