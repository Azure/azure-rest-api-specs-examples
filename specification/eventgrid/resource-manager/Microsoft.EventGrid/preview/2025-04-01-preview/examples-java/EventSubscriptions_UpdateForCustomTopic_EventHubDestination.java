
import com.azure.resourcemanager.eventgrid.models.EventHubEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscription;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import java.util.Arrays;

/**
 * Samples for EventSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_UpdateForCustomTopic_EventHubDestination.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForCustomTopic_EventHubDestination.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForCustomTopicEventHubDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        EventSubscription resource = manager.eventSubscriptions().getWithResponse(
            "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
            "examplesubscription1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDestination(new EventHubEventSubscriptionDestination().withResourceId(
            "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1"))
            .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("existingPrefix")
                .withSubjectEndsWith("newSuffix").withIsSubjectCaseSensitive(true))
            .withLabels(Arrays.asList("label1", "label2")).apply();
    }
}
