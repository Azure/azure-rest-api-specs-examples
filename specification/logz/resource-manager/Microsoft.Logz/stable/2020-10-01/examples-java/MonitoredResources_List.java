import com.azure.core.util.Context;

/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MonitoredResources_List.json
     */
    /**
     * Sample code: MonitoredResources_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitoredResourcesList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", Context.NONE);
    }
}
