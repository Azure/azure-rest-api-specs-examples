/** Samples for MonitoredSubscriptions CreateorUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredSubscriptions_CreateorUpdate.json
     */
    /**
     * Sample code: Monitors_AddMonitoredSubscriptions.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsAddMonitoredSubscriptions(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.monitoredSubscriptions().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
