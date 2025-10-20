
/**
 * Samples for MonitoredSubscriptions CreateorUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/
     * MonitoredSubscriptions_CreateorUpdate.json
     */
    /**
     * Sample code: Monitors_AddMonitoredSubscriptions.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsAddMonitoredSubscriptions(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitoredSubscriptions().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
