/** Samples for SqlResources DeleteSqlRoleAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBSqlRoleAssignmentDelete.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .deleteSqlRoleAssignment(
                "myRoleAssignmentId", "myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
