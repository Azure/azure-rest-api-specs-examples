
import com.azure.resourcemanager.iothub.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2025-08-01-preview/examples/
     * CreateOrReplace_Certificates_With_DeviceRegistryPolicy.json
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
