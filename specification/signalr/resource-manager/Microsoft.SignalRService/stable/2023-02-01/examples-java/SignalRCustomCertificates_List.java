/** Samples for SignalRCustomCertificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalRCustomCertificates_List.json
     */
    /**
     * Sample code: SignalRCustomCertificates_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomCertificatesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomCertificates()
            .list("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
