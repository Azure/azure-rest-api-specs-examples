
import com.azure.resourcemanager.eventgrid.models.AzureFunctionEventSubscriptionDestination;
import com.azure.resourcemanager.eventgrid.models.EventSubscriptionFilter;
import com.azure.resourcemanager.eventgrid.models.StorageBlobDeadLetterDestination;

/**
 * Samples for EventSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
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
            "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1")
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
