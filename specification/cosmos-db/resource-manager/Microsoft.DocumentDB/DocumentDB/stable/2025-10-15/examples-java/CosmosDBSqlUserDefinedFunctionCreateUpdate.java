
import com.azure.resourcemanager.cosmos.models.CreateUpdateOptions;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.SqlUserDefinedFunctionResource;

/**
 * Samples for SqlResources CreateUpdateSqlUserDefinedFunction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBSqlUserDefinedFunctionCreateUpdate.json
     */
    /**
     * Sample code: CosmosDBSqlUserDefinedFunctionCreateUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBSqlUserDefinedFunctionCreateUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getSqlResources().createUpdateSqlUserDefinedFunction("rg1",
            "ddb1", "databaseName", "containerName", "userDefinedFunctionName",
            new SqlUserDefinedFunctionCreateUpdateParameters()
                .withResource(new SqlUserDefinedFunctionResource().withId("userDefinedFunctionName").withBody("body"))
                .withOptions(new CreateUpdateOptions()),
            com.azure.core.util.Context.NONE);
    }
}
