Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlTriggerCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlTriggerResource;
import com.azure.resourcemanager.cosmos.models.TriggerOperation;
import com.azure.resourcemanager.cosmos.models.TriggerType;

/** Samples for SqlResources CreateUpdateSqlTrigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlTriggerCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlTriggerCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlTrigger(
                "rg1",
                "ddb1",
                "databaseName",
                "containerName",
                "triggerName",
                new SqlTriggerCreateUpdateParameters()
                    .withResource(
                        new SqlTriggerResource()
                            .withId("triggerName")
                            .withBody("body")
                            .withTriggerType(TriggerType.fromString("triggerType"))
                            .withTriggerOperation(TriggerOperation.fromString("triggerOperation")))
                    .withOptions(new CreateUpdateOptions()),
                Context.NONE);
    }
}
```
