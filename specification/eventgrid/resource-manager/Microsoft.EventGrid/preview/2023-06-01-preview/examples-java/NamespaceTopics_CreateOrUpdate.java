import com.azure.resourcemanager.eventgrid.models.EventInputSchema;
import com.azure.resourcemanager.eventgrid.models.PublisherType;

/** Samples for NamespaceTopics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopics_CreateOrUpdate.json
     */
    /**
     * Sample code: NamespaceTopics_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .namespaceTopics()
            .define("examplenamespacetopic2")
            .withExistingNamespace("examplerg", "examplenamespace2")
            .withPublisherType(PublisherType.CUSTOM)
            .withInputSchema(EventInputSchema.CLOUD_EVENT_SCHEMA_V1_0)
            .withEventRetentionInDays(1)
            .create();
    }
}
