import com.azure.core.util.Context;

/** Samples for DpsCertificate List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSGetCertificates.json
     */
    /**
     * Sample code: DPSGetCertificates.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetCertificates(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.dpsCertificates().listWithResponse("myResourceGroup", "myFirstProvisioningService", Context.NONE);
    }
}
