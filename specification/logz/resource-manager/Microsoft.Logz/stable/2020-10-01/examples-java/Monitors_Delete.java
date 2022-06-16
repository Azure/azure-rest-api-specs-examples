import com.azure.core.util.Context;

/** Samples for Monitors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Delete.json
     */
    /**
     * Sample code: Monitors_Delete.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitorsDelete(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().delete("myResourceGroup", "myMonitor", Context.NONE);
    }
}
