import com.azure.resourcemanager.datadog.models.MonitoredSubscriptionProperties;

/** Samples for MonitoredSubscriptions Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/MonitoredSubscriptions_Update.json
     */
    /**
     * Sample code: Monitors_UpdateMonitoredSubscriptions.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsUpdateMonitoredSubscriptions(
        com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        MonitoredSubscriptionProperties resource =
            manager
                .monitoredSubscriptions()
                .getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
