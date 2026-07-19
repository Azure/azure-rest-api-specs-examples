
/**
 * Samples for ServerTrustCertificates ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerTrustCertificatesListByInstance.json
     */
    /**
     * Sample code: Gets a list of server trust certificates on a given server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAListOfServerTrustCertificatesOnAGivenServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustCertificates().listByInstance("testrg", "testcl",
            com.azure.core.util.Context.NONE);
    }
}
