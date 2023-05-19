import com.azure.resourcemanager.eventgrid.models.AzureFunctionEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionUpdateParameters;
import com.azure.resourcemanager.eventgrid.models.StorageBlobDeadLetterDestination;

/** Samples for EventSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_UpdateForCustomTopic_AzureFunctionDestination.json
     */
    /**
     * Sample code: EventSubscriptions_UpdateForCustomTopic_AzureFunctionDestination.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsUpdateForCustomTopicAzureFunctionDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .eventSubscriptions()
            .update(
                "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1",
                "examplesubscription1",
                new EventSubscriptionUpdateParameters()
                    .withDestination(
                        new AzureFunctionEventSubscriptionDestination()
                            .withResourceId(
                                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Web/sites/ContosoSite/funtions/ContosoFunc"))
                    .withFilter(
                        new EventSubscriptionFilter()
                            .withSubjectBeginsWith("ExamplePrefix")
                            .withSubjectEndsWith("ExampleSuffix")
                            .withIsSubjectCaseSensitive(false))
                    .withDeadLetterDestination(
                        new StorageBlobDeadLetterDestination()
                            .withResourceId(
                                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg")
                            .withBlobContainerName("contosocontainer")),
                com.azure.core.util.Context.NONE);
    }
}
