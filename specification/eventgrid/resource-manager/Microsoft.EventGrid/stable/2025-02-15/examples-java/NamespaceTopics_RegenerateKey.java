
import com.azure.resourcemanager.eventgrid.models.TopicRegenerateKeyRequest;

/**
 * Samples for NamespaceTopics RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * NamespaceTopics_RegenerateKey.json
     */
    /**
     * Sample code: NamespaceTopics_RegenerateKey.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespaceTopicsRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaceTopics().regenerateKey("examplerg", "examplenamespace2", "examplenamespacetopic2",
            new TopicRegenerateKeyRequest().withKeyName("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
