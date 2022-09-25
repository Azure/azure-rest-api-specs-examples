import com.azure.core.util.Context;

/** Samples for SqlResources ListSqlRoleAssignments. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBSqlRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .listSqlRoleAssignments("myResourceGroupName", "myAccountName", Context.NONE);
    }
}
