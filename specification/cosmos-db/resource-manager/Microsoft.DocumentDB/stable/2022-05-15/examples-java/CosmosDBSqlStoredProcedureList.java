import com.azure.core.util.Context;

/** Samples for SqlResources ListSqlStoredProcedures. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlStoredProcedureList.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlStoredProcedureList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .listSqlStoredProcedures("rgName", "ddb1", "databaseName", "containerName", Context.NONE);
    }
}
