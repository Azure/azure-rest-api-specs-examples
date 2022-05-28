Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.SqlRoleAssignmentCreateUpdateParameters;

/** Samples for SqlResources CreateUpdateSqlRoleAssignment. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlRoleAssignmentCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlRoleAssignmentCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleAssignmentCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlRoleAssignment(
                "myRoleAssignmentId",
                "myResourceGroupName",
                "myAccountName",
                new SqlRoleAssignmentCreateUpdateParameters()
                    .withRoleDefinitionId(
                        "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/sqlRoleDefinitions/myRoleDefinitionId")
                    .withScope(
                        "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases")
                    .withPrincipalId("myPrincipalId"),
                Context.NONE);
    }
}
```
