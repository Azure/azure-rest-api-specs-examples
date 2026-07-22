
/**
 * Samples for SqlResources GetSqlRoleDefinition.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleDefinitionGet.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleDefinitionGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlRoleDefinitionWithResponse("myRoleDefinitionId",
            "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
