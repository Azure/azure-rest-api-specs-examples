import com.azure.core.util.Context;

/** Samples for TrafficFilters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/TrafficFilters_Delete.json
     */
    /**
     * Sample code: TrafficFilters_Delete.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void trafficFiltersDelete(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .trafficFilters()
            .deleteWithResponse("myResourceGroup", "myMonitor", "31d91b5afb6f4c2eaaf104c97b1991dd", Context.NONE);
    }
}
