
/**
 * Samples for ListAssociatedTrafficFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/
     * AssociatedFiltersForDeployment_list.json
     */
    /**
     * Sample code: listAssociatedTrafficFilters_list.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void listAssociatedTrafficFiltersList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.listAssociatedTrafficFilters().listWithResponse("myResourceGroup", "myMonitor",
            com.azure.core.util.Context.NONE);
    }
}
