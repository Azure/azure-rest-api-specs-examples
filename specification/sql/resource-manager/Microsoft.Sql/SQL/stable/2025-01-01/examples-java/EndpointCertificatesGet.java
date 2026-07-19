
/**
 * Samples for EndpointCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EndpointCertificatesGet.json
     */
    /**
     * Sample code: Gets an endpoint certificate.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAnEndpointCertificate(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEndpointCertificates().getWithResponse("testrg", "testcl", "DATABASE_MIRRORING",
            com.azure.core.util.Context.NONE);
    }
}
