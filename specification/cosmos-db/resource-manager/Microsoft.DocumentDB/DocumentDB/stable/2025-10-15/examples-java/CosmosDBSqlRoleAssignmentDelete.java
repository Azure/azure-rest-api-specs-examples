
/**
 * Samples for SqlResources DeleteSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().deleteSqlRoleAssignment(
            "myRoleAssignmentId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
