/** Samples for SignalRCustomDomains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRCustomDomains_List.json
     */
    /**
     * Sample code: SignalRCustomDomains_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomDomainsList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRCustomDomains().list("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
