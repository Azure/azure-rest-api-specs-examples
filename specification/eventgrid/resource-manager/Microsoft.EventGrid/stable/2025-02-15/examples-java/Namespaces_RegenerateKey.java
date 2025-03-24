
import com.azure.resourcemanager.eventgrid.models.NamespaceRegenerateKeyRequest;

/**
 * Samples for Namespaces RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Namespaces_RegenerateKey.
     * json
     */
    /**
     * Sample code: Namespaces_RegenerateKey.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void namespacesRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.namespaces().regenerateKey("examplerg", "exampleNamespaceName1",
            new NamespaceRegenerateKeyRequest().withKeyName("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
