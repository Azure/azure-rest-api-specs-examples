
/**
 * Samples for ServerTrustCertificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerTrustCertificatesDelete.json
     */
    /**
     * Sample code: Delete server trust certificate.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteServerTrustCertificate(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustCertificates().delete("testrg", "testcl", "customerCertificateName",
            com.azure.core.util.Context.NONE);
    }
}
