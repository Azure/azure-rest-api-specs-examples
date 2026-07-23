
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionResource;

/**
 * Samples for SqlResources CreateUpdateSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBSqlUserDefinedFunctionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionCreateUpdate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBSqlUserDefinedFunctionCreateUpdate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getSqlResources().createUpdateSqlUserDefinedFunction("rg1", "ddb1", "databaseName",
            "containerName", "userDefinedFunctionName",
            new SqlUserDefinedFunctionCreateUpdateParameters()
                .withResource(new SqlUserDefinedFunctionResource().withId("userDefinedFunctionName").withBody("body"))
                .withOptions(new CreateUpdateOptions()),
            com.azure.core.util.Context.NONE);
    }
}
