import com.azure.resourcemanager.eventgrid.fluent.models.EventSubscriptionInner;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/** Samples for SystemTopicEventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/SystemTopicEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: SystemTopicEventSubscriptions_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void systemTopicEventSubscriptionsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .systemTopicEventSubscriptions()
            .createOrUpdate(
                "examplerg",
                "exampleSystemTopic1",
                "exampleEventSubscriptionName1",
                new EventSubscriptionInner()
                    .withDestination(
                        new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
                    .withFilter(
                        new EventSubscriptionFilter()
                            .withSubjectBeginsWith("ExamplePrefix")
                            .withSubjectEndsWith("ExampleSuffix")
                            .withIsSubjectCaseSensitive(false)),
                com.azure.core.util.Context.NONE);
    }
}
