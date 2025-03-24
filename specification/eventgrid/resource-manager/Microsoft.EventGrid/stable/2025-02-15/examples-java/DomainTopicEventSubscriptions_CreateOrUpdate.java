
import com.azure.resourcemanager.eventgrid.fluent.models.EventSubscriptionInner;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/**
 * Samples for DomainTopicEventSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * DomainTopicEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: DomainTopicEventSubscriptions_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        domainTopicEventSubscriptionsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainTopicEventSubscriptions().createOrUpdate("examplerg", "exampleDomain1", "exampleDomainTopic1",
            "exampleEventSubscriptionName1",
            new EventSubscriptionInner()
                .withDestination(
                    new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
                .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("ExamplePrefix")
                    .withSubjectEndsWith("ExampleSuffix").withIsSubjectCaseSensitive(false)),
            com.azure.core.util.Context.NONE);
    }
}
