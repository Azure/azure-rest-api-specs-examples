
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionUpdateParameters;
import com.azure.resourcemanager.eventgrid.models.WebhookEventSubscriptionDestination;
import java.util.Arrays;

/**
 * Samples for DomainEventSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * DomainEventSubscriptions_Update.json
     */
    /**
     * Sample code: DomainEventSubscriptions_Update.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainEventSubscriptionsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domainEventSubscriptions().update("examplerg", "exampleDomain1", "exampleEventSubscriptionName1",
            new EventSubscriptionUpdateParameters()
                .withDestination(
                    new WebhookEventSubscriptionDestination().withEndpointUrl("https://requestb.in/15ksip71"))
                .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("existingPrefix")
                    .withSubjectEndsWith("newSuffix").withIsSubjectCaseSensitive(true))
                .withLabels(Arrays.asList("label1", "label2")),
            com.azure.core.util.Context.NONE);
    }
}
