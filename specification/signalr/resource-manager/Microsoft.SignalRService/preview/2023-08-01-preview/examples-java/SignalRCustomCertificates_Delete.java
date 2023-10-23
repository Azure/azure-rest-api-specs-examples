/** Samples for SignalRCustomCertificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRCustomCertificates_Delete.json
     */
    /**
     * Sample code: SignalRCustomCertificates_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomCertificatesDelete(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomCertificates()
            .deleteWithResponse("myResourceGroup", "mySignalRService", "myCert", com.azure.core.util.Context.NONE);
    }
}
