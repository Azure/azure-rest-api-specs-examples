import com.azure.core.util.Context;

/** Samples for ListAssociatedTrafficFilters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/AssociatedFiltersForDeployment_list.json
     */
    /**
     * Sample code: listAssociatedTrafficFilters_list.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void listAssociatedTrafficFiltersList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.listAssociatedTrafficFilters().listWithResponse("myResourceGroup", "myMonitor", Context.NONE);
    }
}
