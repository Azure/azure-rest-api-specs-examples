import com.azure.resourcemanager.eventgrid.fluent.models.EventSubscriptionInner;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/** Samples for DomainEventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/DomainEventSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: DomainEventSubscriptions_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsCreateOrUpdate(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .domainEventSubscriptions()
            .createOrUpdate(
                "examplerg",
                "exampleDomain1",
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
