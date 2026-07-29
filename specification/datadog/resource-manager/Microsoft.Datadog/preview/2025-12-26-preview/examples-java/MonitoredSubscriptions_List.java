
/**
 * Samples for MonitoredSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/MonitoredSubscriptions_List.json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        monitorsGetMonitoredSubscriptions(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitoredSubscriptions().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
