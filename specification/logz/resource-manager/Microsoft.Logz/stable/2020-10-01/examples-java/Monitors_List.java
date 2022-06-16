import com.azure.core.util.Context;

/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_List.json
     */
    /**
     * Sample code: Monitors_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().list(Context.NONE);
    }
}
