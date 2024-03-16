
import com.azure.resourcemanager.newrelicobservability.models.ConfigurationName;

/**
 * Samples for MonitoredSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/
     * MonitoredSubscriptions_Delete.json
     */
    /**
     * Sample code: Monitors_DeleteMonitoredSubscriptions.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsDeleteMonitoredSubscriptions(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitoredSubscriptions().delete("myResourceGroup", "myMonitor", ConfigurationName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
