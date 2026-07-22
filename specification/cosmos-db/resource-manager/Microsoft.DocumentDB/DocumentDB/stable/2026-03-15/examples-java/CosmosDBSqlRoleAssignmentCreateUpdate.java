
import com.azure.resourcemanager.cosmos.models.SqlRoleAssignmentCreateUpdateParameters;

/**
 * Samples for SqlResources CreateUpdateSqlRoleAssignment.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlRoleAssignmentCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlRoleAssignmentCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().createUpdateSqlRoleAssignment("myRoleAssignmentId",
            "myResourceGroupName", "myAccountName",
            new SqlRoleAssignmentCreateUpdateParameters().withRoleDefinitionId(
                "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId")
                .withScope(
                    "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases")
                .withPrincipalId("myPrincipalId"),
            com.azure.core.util.Context.NONE);
    }
}
