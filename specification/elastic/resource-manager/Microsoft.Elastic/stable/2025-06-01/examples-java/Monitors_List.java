
/**
 * Samples for Monitors List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/Monitors_List.json
     */
    /**
     * Sample code: Monitors_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
