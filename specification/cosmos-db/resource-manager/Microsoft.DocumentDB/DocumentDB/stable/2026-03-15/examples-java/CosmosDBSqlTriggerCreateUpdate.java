
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlTriggerCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlTriggerResource;
import com.azure.resourcemanager.cosmos.models.TriggerOperation;
import com.azure.resourcemanager.cosmos.models.TriggerType;

/**
 * Samples for SqlResources CreateUpdateSqlTrigger.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlTriggerCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlTriggerCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBSqlTriggerCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().createUpdateSqlTrigger("rg1", "ddb1", "databaseName", "containerName",
            "triggerName",
            new SqlTriggerCreateUpdateParameters()
                .withResource(new SqlTriggerResource().withId("triggerName").withBody("body")
                    .withTriggerType(TriggerType.fromString("triggerType"))
                    .withTriggerOperation(TriggerOperation.fromString("triggerOperation")))
                .withOptions(new CreateUpdateOptions()),
            com.azure.core.util.Context.NONE);
    }
}
