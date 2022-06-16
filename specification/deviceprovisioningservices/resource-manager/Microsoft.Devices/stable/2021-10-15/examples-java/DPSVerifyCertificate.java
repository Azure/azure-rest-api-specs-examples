import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceprovisioningservices.models.VerificationCodeRequest;

/** Samples for DpsCertificate VerifyCertificate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSVerifyCertificate.json
     */
    /**
     * Sample code: DPSVerifyCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSVerifyCertificate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .verifyCertificateWithResponse(
                "cert",
                "AAAAAAAADGk=",
                "myResourceGroup",
                "myFirstProvisioningService",
                new VerificationCodeRequest().withCertificate("#####################################"),
                null,
                new byte[0],
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
