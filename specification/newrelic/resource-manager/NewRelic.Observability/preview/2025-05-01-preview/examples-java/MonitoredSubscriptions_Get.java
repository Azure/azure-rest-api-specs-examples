
import com.azure.resourcemanager.newrelicobservability.models.ConfigurationName;

/**
 * Samples for MonitoredSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * MonitoredSubscriptions_Get.json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsGetMonitoredSubscriptions(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitoredSubscriptions().getWithResponse("myResourceGroup", "myMonitor", ConfigurationName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
