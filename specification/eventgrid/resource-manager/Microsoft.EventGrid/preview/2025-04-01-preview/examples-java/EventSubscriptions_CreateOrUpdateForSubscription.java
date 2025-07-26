
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;

/**
 * Samples for EventSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_CreateOrUpdateForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_CreateOrUpdateForSubscription.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        eventSubscriptionsCreateOrUpdateForSubscription(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().define("examplesubscription3")
            .withExistingScope("subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40")
            .withDestination(new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
            .withFilter(new EventSubscriptionFilter().withIsSubjectCaseSensitive(false)).create();
    }
}
