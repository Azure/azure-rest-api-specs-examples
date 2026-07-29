
/**
 * Samples for MonitoredSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/MonitoredSubscriptions_Get.json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void
        monitorsGetMonitoredSubscriptions(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitoredSubscriptions().getWithResponse("myResourceGroup", "myMonitor", "default",
            com.azure.core.util.Context.NONE);
    }
}
