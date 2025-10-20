
/**
 * Samples for MonitoredSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/MonitoredSubscriptions_List.
     * json
     */
    /**
     * Sample code: Monitors_GetMonitoredSubscriptions.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsGetMonitoredSubscriptions(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitoredSubscriptions().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
