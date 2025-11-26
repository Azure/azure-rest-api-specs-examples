
import com.azure.resourcemanager.newrelicobservability.models.ConfigurationName;

/**
 * Samples for MonitoredSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * MonitoredSubscriptions_CreateOrUpdate.json
     */
    /**
     * Sample code: Monitors_AddMonitoredSubscriptions.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsAddMonitoredSubscriptions(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitoredSubscriptions().define(ConfigurationName.DEFAULT)
            .withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
