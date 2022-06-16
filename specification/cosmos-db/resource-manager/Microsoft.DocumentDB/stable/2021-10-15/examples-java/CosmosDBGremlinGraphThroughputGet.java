import com.azure.core.util.Context;

/** Samples for GremlinResources GetGremlinGraphThroughput. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBGremlinGraphThroughputGet.json
     */
    /**
     * Sample code: CosmosDBGremlinGraphThroughputGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBGremlinGraphThroughputGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getGremlinResources()
            .getGremlinGraphThroughputWithResponse("rg1", "ddb1", "databaseName", "graphName", Context.NONE);
    }
}
