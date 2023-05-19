import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionUpdateParameters;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;
import java.util.Arrays;

/** Samples for EventSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_UpdateForSubscription.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForSubscription.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForSubscription(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .update(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40",
                "examplesubscription3",
                new EventSubscriptionUpdateParameters()
                    .withDestination(
                        new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
                    .withFilter(
                        new EventSubscriptionFilter()
                            .withSubjectBeginsWith("existingPrefix")
                            .withSubjectEndsWith("newSuffix")
                            .withIsSubjectCaseSensitive(true))
                    .withLabels(Arrays.asList("label1", "label2")),
                com.azure.core.util.Context.NONE);
    }
}
