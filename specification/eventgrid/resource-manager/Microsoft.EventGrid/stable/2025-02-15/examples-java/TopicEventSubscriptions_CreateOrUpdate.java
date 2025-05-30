
import com.azure.resourcemanager.eventgrid.fluent.models.EventSubscriptionInner;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/**
 * Samples for TopicEventSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * TopicEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: TopicEventSubscriptions_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        topicEventSubscriptionsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topicEventSubscriptions().createOrUpdate("examplerg", "exampleTopic1", "exampleEventSubscriptionName1",
            new EventSubscriptionInner()
                .withDestination(
                    new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
                .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("ExamplePrefix")
                    .withSubjectEndsWith("ExampleSuffix").withIsSubjectCaseSensitive(false)),
            com.azure.core.util.Context.NONE);
    }
}
