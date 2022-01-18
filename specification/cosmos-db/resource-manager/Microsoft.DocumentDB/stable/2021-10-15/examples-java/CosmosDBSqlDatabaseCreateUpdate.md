Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlDatabaseCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlDatabaseResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlResources CreateUpdateSqlDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlDatabaseCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlDatabaseCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlDatabaseCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlDatabase(
                "rg1",
                "ddb1",
                "databaseName",
                new SqlDatabaseCreateUpdateParameters()
                    .withLocation("West US")
                    .withTags(mapOf())
                    .withResource(new SqlDatabaseResource().withId("databaseName"))
                    .withOptions(new CreateUpdateOptions()),
                Context.NONE);
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
