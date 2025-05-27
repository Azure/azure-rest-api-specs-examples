
/**
 * Samples for SqlResources ListSqlRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/
     * CosmosDBSqlRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleDefinitionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources()
            .listSqlRoleDefinitions("myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
