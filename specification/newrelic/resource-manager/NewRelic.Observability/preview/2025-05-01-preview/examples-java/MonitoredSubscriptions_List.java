
/**
 * Samples for MonitoredSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * MonitoredSubscriptions_List.json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMonitoredSubscriptions(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitoredSubscriptions().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
