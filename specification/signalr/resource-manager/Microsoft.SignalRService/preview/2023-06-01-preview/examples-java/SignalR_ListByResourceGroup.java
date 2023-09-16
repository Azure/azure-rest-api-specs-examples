/** Samples for SignalR ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalR_ListByResourceGroup.json
     */
    /**
     * Sample code: SignalR_ListByResourceGroup.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRListByResourceGroup(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.signalRs().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
