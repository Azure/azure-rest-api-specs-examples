
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerTrustCertificateInner;

/** Samples for ServerTrustCertificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustCertificatesCreate.json
     */
    /**
     * Sample code: Create server trust certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createServerTrustCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustCertificates().createOrUpdate("testrg", "testcl",
            "customerCertificateName",
            new ServerTrustCertificateInner().withPublicBlob("308203AE30820296A0030201020210"), Context.NONE);
    }
}
