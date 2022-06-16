import com.azure.core.util.Context;

/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/Usages_List.json
     */
    /**
     * Sample code: Usages_List.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void usagesList(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager.usages().list("eastus", Context.NONE);
    }
}
