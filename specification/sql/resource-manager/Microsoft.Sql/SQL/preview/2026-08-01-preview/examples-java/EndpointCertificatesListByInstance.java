
/**
 * Samples for EndpointCertificates ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/EndpointCertificatesListByInstance.json
     */
    /**
     * Sample code: Get a list of endpoint certificates.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAListOfEndpointCertificates(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEndpointCertificates().listByInstance("testrg", "testcl",
            com.azure.core.util.Context.NONE);
    }
}
