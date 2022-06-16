import com.azure.core.util.Context;

/** Samples for Database ListMetricDefinitions. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBDatabaseGetMetricDefinitions.json
     */
    /**
     * Sample code: CosmosDBDatabaseGetMetricDefinitions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBDatabaseGetMetricDefinitions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getDatabases()
            .listMetricDefinitions("rg1", "ddb1", "databaseRid", Context.NONE);
    }
}
