import com.azure.resourcemanager.eventgrid.models.EventHubEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionUpdateParameters;
import java.util.Arrays;

/** Samples for EventSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_UpdateForResourceGroup.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForResourceGroup.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForResourceGroup(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .update(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg",
                "examplesubscription2",
                new EventSubscriptionUpdateParameters()
                    .withDestination(
                        new EventHubEventSubscriptionDestination()
                            .withResourceId(
                                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1"))
                    .withFilter(
                        new EventSubscriptionFilter()
                            .withSubjectBeginsWith("existingPrefix")
                            .withSubjectEndsWith("newSuffix")
                            .withIsSubjectCaseSensitive(true))
                    .withLabels(Arrays.asList("label1", "label2")),
                com.azure.core.util.Context.NONE);
    }
}
