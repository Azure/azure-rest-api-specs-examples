
/**
 * Samples for CaCertificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/
     * CaCertificates_CreateOrUpdate.json
     */
    /**
     * Sample code: CaCertificates_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void caCertificatesCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.caCertificates().define("exampleCACertificateName1")
            .withExistingNamespace("examplerg", "exampleNamespaceName1").withDescription("This is a test certificate")
            .withEncodedCertificate("base64EncodePemFormattedCertificateString").create();
    }
}
