import com.azure.core.util.Context;

/** Samples for SignalRCustomDomains Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_Get.json
     */
    /**
     * Sample code: SignalRCustomDomains_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomDomainsGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRCustomDomains().getWithResponse("myResourceGroup", "mySignalRService", "example", Context.NONE);
    }
}
