
/**
 * Samples for DpsCertificate List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetCertificates.json
     */
    /**
     * Sample code: DPSGetCertificates.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetCertificates(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().list("myResourceGroup", "myFirstProvisioningService",
            com.azure.core.util.Context.NONE);
    }
}
