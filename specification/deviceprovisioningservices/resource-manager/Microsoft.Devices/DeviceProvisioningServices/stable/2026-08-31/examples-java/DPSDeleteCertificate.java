
/**
 * Samples for DpsCertificate Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-31/DPSDeleteCertificate.json
     */
    /**
     * Sample code: DPSDeleteCertificate.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSDeleteCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().deleteWithResponse("myResourceGroup", "AAAAAAAADGk=", "myFirstProvisioningService",
            "cert", null, null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
