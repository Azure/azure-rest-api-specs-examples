/** Samples for MonitoredSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredSubscriptions_Delete.json
     */
    /**
     * Sample code: Monitors_DeleteMonitoredSubscriptions.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsDeleteMonitoredSubscriptions(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager
            .monitoredSubscriptions()
            .delete("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
