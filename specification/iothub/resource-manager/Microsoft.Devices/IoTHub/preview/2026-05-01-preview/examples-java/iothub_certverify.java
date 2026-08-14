
import com.azure.resourcemanager.iothub.models.CertificateVerificationDescription;

/**
 * Samples for Certificates Verify.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_certverify.json
     */
    /**
     * Sample code: Certificates_Verify.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesVerify(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().verifyWithResponse("myResourceGroup", "myFirstProvisioningService", "cert",
            "AAAAAAAADGk=",
            new CertificateVerificationDescription().withCertificate("#####################################"),
            com.azure.core.util.Context.NONE);
    }
}
