/** Samples for SignalRCustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRCustomDomains_Get.json
     */
    /**
     * Sample code: SignalRCustomDomains_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomDomainsGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRCustomDomains()
            .getWithResponse("myResourceGroup", "mySignalRService", "example", com.azure.core.util.Context.NONE);
    }
}
