
import com.azure.resourcemanager.deviceprovisioningservices.models.CertificateProperties;

/**
 * Samples for DpsCertificate CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSCertificateCreateOrUpdate.json
     */
    /**
     * Sample code: DPSCreateOrUpdateCertificate.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSCreateOrUpdateCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().define("cert")
            .withExistingProvisioningService("myResourceGroup", "myFirstProvisioningService")
            .withProperties(new CertificateProperties().withCertificate("MA==".getBytes())).create();
    }
}
