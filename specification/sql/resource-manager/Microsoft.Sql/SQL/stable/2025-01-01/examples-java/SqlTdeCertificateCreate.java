
import com.azure.resourcemanager.sql.models.TdeCertificate;

/**
 * Samples for TdeCertificates Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SqlTdeCertificateCreate.json
     */
    /**
     * Sample code: Upload a TDE certificate.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void uploadATDECertificate(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTdeCertificates().create("testtdecert", "testtdecert",
            new TdeCertificate().withPrivateBlob("MIIXXXXXXXX"), com.azure.core.util.Context.NONE);
    }
}
