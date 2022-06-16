import com.azure.core.util.Context;

/** Samples for SqlResources GetSqlContainer. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlContainerGet.json
     */
    /**
     * Sample code: CosmosDBSqlContainerGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlContainerGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .getSqlContainerWithResponse("rgName", "ddb1", "databaseName", "containerName", Context.NONE);
    }
}
