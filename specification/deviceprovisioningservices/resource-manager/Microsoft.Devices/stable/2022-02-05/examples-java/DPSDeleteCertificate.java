import com.azure.core.util.Context;

/** Samples for DpsCertificate Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSDeleteCertificate.json
     */
    /**
     * Sample code: DPSDeleteCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSDeleteCertificate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .deleteWithResponse(
                "myResourceGroup",
                "AAAAAAAADGk=",
                "myFirstProvisioningService",
                "cert",
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
