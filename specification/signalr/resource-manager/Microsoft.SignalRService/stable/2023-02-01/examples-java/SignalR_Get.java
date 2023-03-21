/** Samples for SignalR GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/SignalR_Get.json
     */
    /**
     * Sample code: SignalR_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRs()
            .getByResourceGroupWithResponse("myResourceGroup", "mySignalRService", com.azure.core.util.Context.NONE);
    }
}
