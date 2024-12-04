
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionUpdateParameters;
import com.azure.resourcemanager.eventgrid.models.ServiceBusTopicEventSubscriptionDestination;
import java.util.Arrays;

/**
 * Samples for EventSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * EventSubscriptions_UpdateForCustomTopic_ServiceBusTopicDestination.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForCustomTopic_ServiceBusTopicDestination.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForCustomTopicServiceBusTopicDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().update(
            "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2",
            "examplesubscription1",
            new EventSubscriptionUpdateParameters()
                .withDestination(new ServiceBusTopicEventSubscriptionDestination().withResourceId(
                    "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.ServiceBus/namespaces/ContosoNamespace/topics/SBT"))
                .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("existingPrefix")
                    .withSubjectEndsWith("newSuffix").withIsSubjectCaseSensitive(true))
                .withLabels(Arrays.asList("label1", "label2")),
            com.azure.core.util.Context.NONE);
    }
}
