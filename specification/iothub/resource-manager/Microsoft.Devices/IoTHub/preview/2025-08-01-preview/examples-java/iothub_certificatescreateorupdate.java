
import com.azure.resourcemanager.iothub.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2025-08-01-preview/examples/
     * iothub_certificatescreateorupdate.json
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
