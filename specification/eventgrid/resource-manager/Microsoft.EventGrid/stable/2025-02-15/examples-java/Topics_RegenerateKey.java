
import com.azure.resourcemanager.eventgrid.models.TopicRegenerateKeyRequest;

/**
 * Samples for Topics RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Topics_RegenerateKey.json
     */
    /**
     * Sample code: Topics_RegenerateKey.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void topicsRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.topics().regenerateKey("examplerg", "exampletopic2",
            new TopicRegenerateKeyRequest().withKeyName("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
