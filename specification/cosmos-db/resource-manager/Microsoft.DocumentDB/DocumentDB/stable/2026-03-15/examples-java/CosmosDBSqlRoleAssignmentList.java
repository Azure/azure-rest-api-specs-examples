
/**
 * Samples for SqlResources ListSqlRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleAssignmentList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().listSqlRoleAssignments("myResourceGroupName", "myAccountName",
            com.azure.core.util.Context.NONE);
    }
}
