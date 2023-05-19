import com.azure.resourcemanager.eventgrid.fluent.models.EventSubscriptionInner;
import com.azure.resourcemanager.eventgrid.models.EventHubEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;

/** Samples for EventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_WebhookDestination.json
     */
    /**
     * Sample code: EventSubscriptions_CreateOrUpdateForCustomTopic_WebhookDestination.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsCreateOrUpdateForCustomTopicWebhookDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .createOrUpdate(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
                "examplesubscription1",
                new EventSubscriptionInner()
                    .withDestination(
                        new EventHubEventSubscriptionDestination()
                            .withResourceId(
                                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1"))
                    .withFilter(
                        new EventSubscriptionFilter()
                            .withSubjectBeginsWith("ExamplePrefix")
                            .withSubjectEndsWith("ExampleSuffix")
                            .withIsSubjectCaseSensitive(false)),
                com.azure.core.util.Context.NONE);
    }
}
