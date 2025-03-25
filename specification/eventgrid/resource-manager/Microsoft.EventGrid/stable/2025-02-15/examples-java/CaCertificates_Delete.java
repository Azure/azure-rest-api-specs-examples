
/**
 * Samples for CaCertificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/CaCertificates_Delete.
     * json
     */
    /**
     * Sample code: CaCertificates_Delete.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void caCertificatesDelete(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.caCertificates().delete("examplerg", "exampleNamespaceName1", "exampleCACertificateName1",
            com.azure.core.util.Context.NONE);
    }
}
