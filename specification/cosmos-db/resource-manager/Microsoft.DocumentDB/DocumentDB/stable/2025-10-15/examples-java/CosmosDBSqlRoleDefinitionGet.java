
/**
 * Samples for SqlResources GetSqlRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleDefinitionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlRoleDefinitionWithResponse(
            "myRoleDefinitionId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
