import com.azure.core.util.Context;

/** Samples for DpsCertificate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetCertificate.json
     */
    /**
     * Sample code: DPSGetCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .getWithResponse("cert", "myResourceGroup", "myFirstProvisioningService", null, Context.NONE);
    }
}
