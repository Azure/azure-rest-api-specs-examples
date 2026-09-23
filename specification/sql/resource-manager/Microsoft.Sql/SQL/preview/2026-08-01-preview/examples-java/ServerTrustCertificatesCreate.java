
import com.azure.resourcemanager.sql.fluent.models.ServerTrustCertificateInner;

/**
 * Samples for ServerTrustCertificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerTrustCertificatesCreate.json
     */
    /**
     * Sample code: Create server trust certificate.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createServerTrustCertificate(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerTrustCertificates().createOrUpdate("testrg", "testcl",
            "customerCertificateName",
            new ServerTrustCertificateInner().withPublicBlob("308203AE30820296A0030201020210"),
            com.azure.core.util.Context.NONE);
    }
}
