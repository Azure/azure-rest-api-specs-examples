
import com.azure.resourcemanager.newrelicobservability.models.ConfigurationName;
import com.azure.resourcemanager.newrelicobservability.models.MonitoredSubscriptionProperties;

/**
 * Samples for MonitoredSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-01-01/examples/
     * MonitoredSubscriptions_Update.json
     */
    /**
     * Sample code: Monitors_UpdateMonitoredSubscriptions.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsUpdateMonitoredSubscriptions(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        MonitoredSubscriptionProperties resource = manager.monitoredSubscriptions().getWithResponse("myResourceGroup",
            "myMonitor", ConfigurationName.DEFAULT, com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
