
/**
 * Samples for SqlResources DeleteSqlRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleDefinitionDelete.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleDefinitionDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlRoleDefinition("myRoleDefinitionId", "myResourceGroupName",
            "myAccountName", com.azure.core.util.Context.NONE);
    }
}
