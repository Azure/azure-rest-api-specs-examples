
/**
 * Samples for Monitors ListMonitoredResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/MonitoredResources_List.json
     */
    /**
     * Sample code: Monitors_ListMonitoredResources.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        monitorsListMonitoredResources(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
