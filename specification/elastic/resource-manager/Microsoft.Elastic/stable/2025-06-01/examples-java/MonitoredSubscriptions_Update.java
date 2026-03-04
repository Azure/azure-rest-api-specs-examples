
import com.azure.resourcemanager.elastic.models.MonitoredSubscriptionProperties;

/**
 * Samples for MonitoredSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/MonitoredSubscriptions_Update.json
     */
    /**
     * Sample code: Monitors_UpdateMonitoredSubscriptions.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsUpdateMonitoredSubscriptions(com.azure.resourcemanager.elastic.ElasticManager manager) {
        MonitoredSubscriptionProperties resource = manager.monitoredSubscriptions()
            .getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
