
import com.azure.resourcemanager.eventgrid.models.PartnerNamespaceRegenerateKeyRequest;

/**
 * Samples for PartnerNamespaces RegenerateKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * PartnerNamespaces_RegenerateKey.json
     */
    /**
     * Sample code: PartnerNamespaces_RegenerateKey.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerNamespacesRegenerateKey(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.partnerNamespaces().regenerateKeyWithResponse("examplerg", "examplePartnerNamespaceName1",
            new PartnerNamespaceRegenerateKeyRequest().withKeyName("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
