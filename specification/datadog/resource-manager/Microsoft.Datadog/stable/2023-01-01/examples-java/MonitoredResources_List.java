/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredResources_List.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsListMonitoredResources(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
