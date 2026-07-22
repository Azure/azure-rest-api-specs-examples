
/**
 * Samples for SqlResources ListSqlRoleDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleDefinitionList.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleDefinitionList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlRoleDefinitions("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
