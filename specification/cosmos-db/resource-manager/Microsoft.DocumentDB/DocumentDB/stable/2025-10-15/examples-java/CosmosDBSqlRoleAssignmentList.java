
/**
 * Samples for SqlResources ListSqlRoleAssignments.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlRoleAssignmentList.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources()
            .listSqlRoleAssignments("myResourceGroupName", "myAccountName", com.azure.core.util.Context.NONE);
    }
}
