import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/** Samples for EventSubscriptions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_CreateOrUpdateForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_CreateOrUpdateForSubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsCreateOrUpdateForSubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .define("examplesubscription3")
            .withExistingScope("subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4")
            .withDestination(new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
            .withFilter(new EventSubscriptionFilter().withIsSubjectCaseSensitive(false))
            .create();
    }
}
