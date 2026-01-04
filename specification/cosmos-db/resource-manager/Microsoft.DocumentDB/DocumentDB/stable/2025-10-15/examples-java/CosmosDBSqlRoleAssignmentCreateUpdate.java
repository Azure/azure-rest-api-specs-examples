
import com.azure.resourcemanager.cosmos.models.SqlRoleAssignmentCreateUpdateParameters;

/**
 * Samples for SqlResources CreateUpdateSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlRoleAssignmentCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().createUpdateSqlRoleAssignment(
            "myRoleAssignmentId", "myResourceGroupName", "myAccountName",
            new SqlRoleAssignmentCreateUpdateParameters().withRoleDefinitionId(
                "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId")
                .withScope(
                    "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases")
                .withPrincipalId("myPrincipalId"),
            com.azure.core.util.Context.NONE);
    }
}
