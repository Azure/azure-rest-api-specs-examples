import com.azure.core.util.Context;

/** Samples for Certificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_getcertificate.json
     */
    /**
     * Sample code: Certificates_Get.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesGet(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().getWithResponse("myResourceGroup", "testhub", "cert", Context.NONE);
    }
}
