/** Samples for DpsCertificate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSCertificateCreateOrUpdate.json
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
            .withCertificate("############################################")
            .create();
    }
}
