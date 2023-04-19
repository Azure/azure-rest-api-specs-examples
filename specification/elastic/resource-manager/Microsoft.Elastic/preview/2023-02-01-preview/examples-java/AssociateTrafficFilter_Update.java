/** Samples for AssociateTrafficFilter Associate. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/AssociateTrafficFilter_Update.json
     */
    /**
     * Sample code: AssociateTrafficFilter_Associate.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void associateTrafficFilterAssociate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .associateTrafficFilters()
            .associate(
                "myResourceGroup", "myMonitor", "31d91b5afb6f4c2eaaf104c97b1991dd", com.azure.core.util.Context.NONE);
    }
}
