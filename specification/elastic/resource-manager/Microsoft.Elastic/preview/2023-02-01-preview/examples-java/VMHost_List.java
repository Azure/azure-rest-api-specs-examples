/** Samples for VMHost List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/VMHost_List.json
     */
    /**
     * Sample code: VMHost_List.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void vMHostList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.vMHosts().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
