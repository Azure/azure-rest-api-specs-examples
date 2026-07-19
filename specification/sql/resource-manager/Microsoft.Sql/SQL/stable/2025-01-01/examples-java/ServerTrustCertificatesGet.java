
/**
 * Samples for ServerTrustCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerTrustCertificatesGet.json
     */
    /**
     * Sample code: Gets server trust certificate.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsServerTrustCertificate(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustCertificates().getWithResponse("testrg", "testcl",
            "customerCertificateName", com.azure.core.util.Context.NONE);
    }
}
