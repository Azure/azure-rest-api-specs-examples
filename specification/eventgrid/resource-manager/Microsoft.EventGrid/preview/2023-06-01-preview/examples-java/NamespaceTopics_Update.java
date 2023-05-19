import com.azure.resourcemanager.eventgrid.models.NamespaceTopic;

/** Samples for NamespaceTopics Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopics_Update.json
     */
    /**
     * Sample code: NamespaceTopics_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        NamespaceTopic resource =
            manager
                .namespaceTopics()
                .getWithResponse(
                    "examplerg",
                    "exampleNamespaceName1",
                    "exampleNamespaceTopicName1",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withEventRetentionInDays(1).apply();
    }
}
