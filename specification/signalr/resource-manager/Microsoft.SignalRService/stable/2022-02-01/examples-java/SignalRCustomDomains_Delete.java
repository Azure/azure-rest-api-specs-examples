import com.azure.core.util.Context;

/** Samples for SignalRCustomDomains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomDomains_Delete.json
     */
    /**
     * Sample code: SignalRCustomDomains_Delete.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRCustomDomainsDelete(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRCustomDomains().delete("myResourceGroup", "mySignalRService", "example", Context.NONE);
    }
}
