/** Samples for MonitorOperation Upgrade. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/Monitor_Upgrade.json
     */
    /**
     * Sample code: Monitor_Upgrade.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorUpgrade(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitorOperations().upgrade("myResourceGroup", "myMonitor", null, com.azure.core.util.Context.NONE);
    }
}
