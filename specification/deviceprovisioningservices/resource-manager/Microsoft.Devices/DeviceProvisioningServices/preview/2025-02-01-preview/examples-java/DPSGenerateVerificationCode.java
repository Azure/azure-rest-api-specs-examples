
/**
 * Samples for DpsCertificate GenerateVerificationCode.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGenerateVerificationCode.json
     */
    /**
     * Sample code: DPSGenerateVerificationCode.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSGenerateVerificationCode(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().generateVerificationCodeWithResponse("cert", "AAAAAAAADGk=", "myResourceGroup",
            "myFirstProvisioningService", null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
