import com.azure.resourcemanager.deviceprovisioningservices.models.CertificateProperties;

/** Samples for DpsCertificate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCertificateCreateOrUpdate.json
     */
    /**
     * Sample code: DPSCreateOrUpdateCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSCreateOrUpdateCertificate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .define("cert")
            .withExistingProvisioningService("myResourceGroup", "myFirstProvisioningService")
            .withProperties(
                new CertificateProperties().withCertificate("############################################".getBytes()))
            .create();
    }
}
