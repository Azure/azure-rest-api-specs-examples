Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlResources ListSqlRoleAssignments. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlRoleAssignmentList.json
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
```
