
/**
 * Samples for SqlResources GetSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/
     * CosmosDBSqlRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().getSqlRoleAssignmentWithResponse(
            "myRoleAssignmentId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
