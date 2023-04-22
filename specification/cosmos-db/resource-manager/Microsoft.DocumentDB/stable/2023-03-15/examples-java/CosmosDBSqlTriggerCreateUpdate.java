import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlTriggerCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlTriggerResource;
import com.azure.resourcemanager.cosmos.models.TriggerOperation;
import com.azure.resourcemanager.cosmos.models.TriggerType;

/** Samples for SqlResources CreateUpdateSqlTrigger. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBSqlTriggerCreateUpdate.json
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
                com.azure.core.util.Context.NONE);
    }
}
