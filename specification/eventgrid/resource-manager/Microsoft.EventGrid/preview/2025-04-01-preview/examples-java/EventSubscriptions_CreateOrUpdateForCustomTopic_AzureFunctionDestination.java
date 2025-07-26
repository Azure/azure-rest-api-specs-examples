
import com.azure.resourcemanager.eventgrid.models.AzureFunctionEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.StorageBlobDeadLetterDestination;

/**
 * Samples for EventSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * EventSubscriptions_CreateOrUpdateForCustomTopic_AzureFunctionDestination.json
     */
    /**
     * Sample code: EventSubscriptions_CreateOrUpdateForCustomTopic_AzureFunctionDestination.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void eventSubscriptionsCreateOrUpdateForCustomTopicAzureFunctionDestination(
        com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.eventSubscriptions().define("examplesubscription1").withExistingScope(
            "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1")
            .withDestination(new AzureFunctionEventSubscriptionDestination().withResourceId(
                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Web/sites/ContosoSite/funtions/ContosoFunc"))
            .withFilter(new EventSubscriptionFilter().withSubjectBeginsWith("ExamplePrefix")
                .withSubjectEndsWith("ExampleSuffix").withIsSubjectCaseSensitive(false))
            .withDeadLetterDestination(new StorageBlobDeadLetterDestination().withResourceId(
                "/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg")
                .withBlobContainerName("contosocontainer"))
            .create();
    }
}
