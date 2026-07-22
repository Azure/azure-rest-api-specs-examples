
/**
 * Samples for SqlResources GetSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentGet.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleAssignmentGet(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().getSqlRoleAssignmentWithResponse("myRoleAssignmentId",
            "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
