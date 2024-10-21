
/**
 * Samples for AllTrafficFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/AllTrafficFilters_list.json
     */
    /**
     * Sample code: AllTrafficFilters_list.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void allTrafficFiltersList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.allTrafficFilters().listWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
