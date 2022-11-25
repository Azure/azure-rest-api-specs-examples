import com.azure.core.util.Context;

/** Samples for MonitoredResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/MonitoredResources_List.json
     */
    /**
     * Sample code: MonitoredResources_List.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void monitoredResourcesList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitoredResources().list("myResourceGroup", "myMonitor", Context.NONE);
    }
}
