/** Samples for MonitoredSubscriptions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredSubscriptions_List.json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsGetMonitoredSubscriptions(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitoredSubscriptions().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
