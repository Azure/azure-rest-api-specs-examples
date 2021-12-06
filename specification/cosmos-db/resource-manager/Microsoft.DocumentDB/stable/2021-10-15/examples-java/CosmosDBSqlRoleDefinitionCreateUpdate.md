Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.Permission;
import com.azure.resourcemanager.cosmos.models.RoleDefinitionType;
import com.azure.resourcemanager.cosmos.models.SqlRoleDefinitionCreateUpdateParameters;
import java.util.Arrays;

/** Samples for SqlResources CreateUpdateSqlRoleDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlRoleDefinitionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlRoleDefinitionCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlRoleDefinitionCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlRoleDefinition(
                "myRoleDefinitionId",
                "myResourceGroupName",
                "myAccountName",
                new SqlRoleDefinitionCreateUpdateParameters()
                    .withRoleName("myRoleName")
                    .withType(RoleDefinitionType.CUSTOM_ROLE)
                    .withAssignableScopes(
                        Arrays
                            .asList(
                                "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/sales",
                                "/subscriptions/mySubscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases"))
                    .withPermissions(
                        Arrays
                            .asList(
                                new Permission()
                                    .withDataActions(
                                        Arrays
                                            .asList(
                                                "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/create",
                                                "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/items/read"))
                                    .withNotDataActions(Arrays.asList()))),
                Context.NONE);
    }
}
```
