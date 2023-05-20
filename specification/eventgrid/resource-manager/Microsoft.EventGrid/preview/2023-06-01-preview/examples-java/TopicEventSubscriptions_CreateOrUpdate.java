import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/** Samples for TopicEventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/TopicEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: TopicEventSubscriptions_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void topicEventSubscriptionsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .topicEventSubscriptions()
            .define("exampleEventSubscriptionName1")
            .withExistingTopic("examplerg", "exampleTopic1")
            .withDestination(new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
            .withFilter(
                new EventSubscriptionFilter()
                    .withSubjectBeginsWith("ExamplePrefix")
                    .withSubjectEndsWith("ExampleSuffix")
                    .withIsSubjectCaseSensitive(false))
            .create();
    }
}
