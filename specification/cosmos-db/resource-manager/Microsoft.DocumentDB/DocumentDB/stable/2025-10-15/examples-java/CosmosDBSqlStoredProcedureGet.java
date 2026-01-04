
/**
 * Samples for SqlResources GetSqlStoredProcedure.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlStoredProcedureGet.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlStoredProcedureGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlStoredProcedureWithResponse("rgName",
            "ddb1", "databaseName", "containerName", "storedProcedureName", com.azure.core.util.Context.NONE);
    }
}
