
import com.azure.resourcemanager.iothub.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/iothub_certificatescreateorupdate.json
     */
    /**
     * Sample code: Certificates_CreateOrUpdate.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesCreateOrUpdate(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().define("cert").withExistingIotHub("myResourceGroup", "iothub")
            .withProperties(new CertificateProperties().withCertificate("############################################"))
            .create();
    }
}
