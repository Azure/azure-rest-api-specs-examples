
/**
 * Samples for Collection ListMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBCollectionGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBCollectionGetMetricDefinitions.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBCollectionGetMetricDefinitions(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCollections().listMetricDefinitions("rg1", "ddb1", "databaseRid", "collectionRid",
            com.azure.core.util.Context.NONE);
    }
}
