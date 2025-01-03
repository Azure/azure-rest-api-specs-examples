
import com.azure.resourcemanager.deviceprovisioningservices.models.VerificationCodeRequest;

/**
 * Samples for DpsCertificate VerifyCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/
     * DPSVerifyCertificate.json
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
