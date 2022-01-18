Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlStoredProcedureCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlStoredProcedureResource;

/** Samples for SqlResources CreateUpdateSqlStoredProcedure. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBSqlStoredProcedureCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlStoredProcedureCreateUpdate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBSqlStoredProcedureCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getSqlResources()
            .createUpdateSqlStoredProcedure(
                "rg1",
                "ddb1",
                "databaseName",
                "containerName",
                "storedProcedureName",
                new SqlStoredProcedureCreateUpdateParameters()
                    .withResource(new SqlStoredProcedureResource().withId("storedProcedureName").withBody("body"))
                    .withOptions(new CreateUpdateOptions()),
                Context.NONE);
    }
}
```
