/** Samples for SignalRReplicas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/SignalRReplicas_Get.json
     */
    /**
     * Sample code: SignalRReplicas_Get.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRReplicasGet(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRReplicas()
            .getWithResponse(
                "myResourceGroup", "mySignalRService", "mySignalRService-eastus", com.azure.core.util.Context.NONE);
    }
}
