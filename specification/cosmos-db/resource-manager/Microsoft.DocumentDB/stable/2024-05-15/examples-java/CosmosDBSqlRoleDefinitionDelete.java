
/**
 * Samples for SqlResources DeleteSqlRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBSqlRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleDefinitionDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlRoleDefinition(
            "myRoleDefinitionId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
