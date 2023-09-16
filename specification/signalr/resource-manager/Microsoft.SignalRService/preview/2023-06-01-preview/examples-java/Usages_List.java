/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/Usages_List.json
     */
    /**
     * Sample code: Usages_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void usagesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.usages().list("eastus", com.azure.core.util.Context.NONE);
    }
}
