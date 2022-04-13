Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionResource;

/** Samples for SqlResources CreateUpdateSqlUserDefinedFunction. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlUserDefinedFunctionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlUserDefinedFunctionCreateUpdate(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlUserDefinedFunction(
                "rg1",
                "ddb1",
                "databaseName",
                "containerName",
                "userDefinedFunctionName",
                new SqlUserDefinedFunctionCreateUpdateParameters()
                    .withResource(
                        new SqlUserDefinedFunctionResource().withId("userDefinedFunctionName").withBody("body"))
                    .withOptions(new CreateUpdateOptions()),
                Context.NONE);
    }
}
```
