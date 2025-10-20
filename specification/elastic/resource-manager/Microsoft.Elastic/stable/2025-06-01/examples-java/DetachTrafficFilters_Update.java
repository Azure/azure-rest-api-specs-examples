
/**
 * Samples for DetachTrafficFilter Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/DetachTrafficFilters_Update.
     * json
     */
    /**
     * Sample code: DetachTrafficFilter_Update.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void detachTrafficFilterUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.detachTrafficFilters().update("myResourceGroup", "myMonitor", "31d91b5afb6f4c2eaaf104c97b1991dd",
            com.azure.core.util.Context.NONE);
    }
}
