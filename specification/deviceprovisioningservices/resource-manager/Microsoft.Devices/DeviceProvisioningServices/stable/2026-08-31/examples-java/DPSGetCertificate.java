
/**
 * Samples for DpsCertificate Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-31/DPSGetCertificate.json
     */
    /**
     * Sample code: DPSGetCertificate.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().getWithResponse("cert", "myResourceGroup", "myFirstProvisioningService", null,
            com.azure.core.util.Context.NONE);
    }
}
