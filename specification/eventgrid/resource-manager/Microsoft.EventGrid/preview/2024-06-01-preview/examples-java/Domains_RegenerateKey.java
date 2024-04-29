
import com.azure.resourcemanager.eventgrid.models.DomainRegenerateKeyRequest;

/**
 * Samples for Domains RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * Domains_RegenerateKey.json
     */
    /**
     * Sample code: Domains_RegenerateKey.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void domainsRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.domains().regenerateKeyWithResponse("examplerg", "exampledomain2",
            new DomainRegenerateKeyRequest().withKeyName("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
