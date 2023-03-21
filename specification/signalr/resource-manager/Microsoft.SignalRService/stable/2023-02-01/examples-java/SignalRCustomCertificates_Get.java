/** Samples for SignalRCustomCertificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalRCustomCertificates_Get.json
     */
    /**
     * Sample code: SignalRCustomCertificates_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomCertificatesGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomCertificates()
            .getWithResponse("myResourceGroup", "mySignalRService", "myCert", com.azure.core.util.Context.NONE);
    }
}
