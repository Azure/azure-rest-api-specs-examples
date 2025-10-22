
import com.azure.resourcemanager.deviceprovisioningservices.models.VerificationCodeRequest;

/**
 * Samples for DpsCertificate VerifyCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSVerifyCertificate.json
     */
    /**
     * Sample code: DPSVerifyCertificate.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        dPSVerifyCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().verifyCertificateWithResponse("cert", "AAAAAAAADGk=", "myResourceGroup",
            "myFirstProvisioningService",
            new VerificationCodeRequest().withCertificate("#####################################"), null, null, null,
            null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
