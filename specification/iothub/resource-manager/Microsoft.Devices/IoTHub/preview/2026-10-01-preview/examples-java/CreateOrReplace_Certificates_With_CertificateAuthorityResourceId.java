
import com.azure.resourcemanager.iothub.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/CreateOrReplace_Certificates_With_CertificateAuthorityResourceId.json
     */
    /**
     * Sample code: CreateOrReplace_Certificates_With_CertificateAuthorityResourceId.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void createOrReplaceCertificatesWithCertificateAuthorityResourceId(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().define("cert").withExistingIotHub("myResourceGroup", "testHub")
            .withProperties(new CertificateProperties().withCertificate("############################################"))
            .create();
    }
}
