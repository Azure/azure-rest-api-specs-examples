import com.azure.core.util.Context;

/** Samples for SignalRCustomCertificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_Get.json
     */
    /**
     * Sample code: SignalRCustomCertificates_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomCertificatesGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomCertificates()
            .getWithResponse("myResourceGroup", "mySignalRService", "myCert", Context.NONE);
    }
}
