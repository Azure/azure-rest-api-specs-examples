
/**
 * Samples for SqlResources DeleteSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleAssignmentDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().deleteSqlRoleAssignment("myRoleAssignmentId", "myResourceGroupName",
            "myAccountName", com.azure.core.util.Context.NONE);
    }
}
