import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void operationsList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.operations().list(Context.NONE);
    }
}
