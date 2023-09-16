/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-06-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void operationsList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
