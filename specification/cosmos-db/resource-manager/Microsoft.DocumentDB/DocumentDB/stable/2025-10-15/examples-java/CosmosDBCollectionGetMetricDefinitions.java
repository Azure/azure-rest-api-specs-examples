
/**
 * Samples for Collection ListMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBCollectionGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBCollectionGetMetricDefinitions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBCollectionGetMetricDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getCollections().listMetricDefinitions("rg1", "ddb1",
            "databaseRid", "collectionRid", com.azure.core.util.Context.NONE);
    }
}
