
import com.azure.resourcemanager.iothub.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/CreateOrReplace_Certificates_With_DeviceRegistryPolicy.json
     */
    /**
     * Sample code: CreateOrReplace_Certificates_With_DeviceRegistryPolicy.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void
        createOrReplaceCertificatesWithDeviceRegistryPolicy(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().define("cert").withExistingIotHub("myResourceGroup", "testHub")
            .withProperties(new CertificateProperties().withCertificate("############################################"))
            .create();
    }
}
