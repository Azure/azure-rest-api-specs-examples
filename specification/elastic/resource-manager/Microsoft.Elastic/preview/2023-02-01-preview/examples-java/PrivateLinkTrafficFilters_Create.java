/** Samples for CreateAndAssociatePLFilter Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/PrivateLinkTrafficFilters_Create.json
     */
    /**
     * Sample code: createAndAssociatePLFilter_Create.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void createAndAssociatePLFilterCreate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .createAndAssociatePLFilters()
            .create(
                "myResourceGroup",
                "myMonitor",
                null,
                "fdb54d3b-e85e-4d08-8958-0d2f7g523df9",
                "myPrivateEndpoint",
                com.azure.core.util.Context.NONE);
    }
}
