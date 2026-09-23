
/**
 * Samples for EndpointCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/EndpointCertificatesGetWithTrustedRootCertificate.json
     */
    /**
     * Sample code: Gets an endpoint certificate including trusted root certificates.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAnEndpointCertificateIncludingTrustedRootCertificates(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEndpointCertificates().getWithResponse("testrg", "testcl", "DATABASE_MIRRORING",
            com.azure.core.util.Context.NONE);
    }
}
