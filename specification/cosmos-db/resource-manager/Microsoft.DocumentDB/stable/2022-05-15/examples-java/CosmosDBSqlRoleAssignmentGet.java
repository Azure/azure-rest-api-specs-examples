import com.azure.core.util.Context;

/** Samples for SqlResources GetSqlRoleAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBSqlRoleAssignmentGet.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .getSqlRoleAssignmentWithResponse(
                "myRoleAssignmentId", "myResourceGroupName", "myAccountName", Context.NONE);
    }
}
