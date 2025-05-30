
/**
 * Samples for CaCertificates ListByNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * CaCertificates_ListByNamespace.json
     */
    /**
     * Sample code: CaCertificates_ListByNamespace.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void caCertificatesListByNamespace(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.caCertificates().listByNamespace("examplerg", "namespace123", null, null,
            com.azure.core.util.Context.NONE);
    }
}
