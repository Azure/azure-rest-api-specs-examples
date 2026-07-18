
/**
 * Samples for EndpointCertificates ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EndpointCertificatesListByInstance.json
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
