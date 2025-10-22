
/**
 * Samples for DpsCertificate Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetCertificate.json
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
