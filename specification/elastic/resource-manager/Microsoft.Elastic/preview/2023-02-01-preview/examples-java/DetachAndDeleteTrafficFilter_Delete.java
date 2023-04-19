/** Samples for DetachAndDeleteTrafficFilter Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/DetachAndDeleteTrafficFilter_Delete.json
     */
    /**
     * Sample code: DetachAndDeleteTrafficFilter_Delete.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void detachAndDeleteTrafficFilterDelete(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .detachAndDeleteTrafficFilters()
            .deleteWithResponse(
                "myResourceGroup", "myMonitor", "31d91b5afb6f4c2eaaf104c97b1991dd", com.azure.core.util.Context.NONE);
    }
}
